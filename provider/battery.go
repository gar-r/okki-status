package provider

import (
	"fmt"
	"okki-status/core"

	"github.com/godbus/dbus/v5"
)

// BatteryUpdate contains extra attributes according to the
// UPower dbus interface: https://upower.freedesktop.org/docs/Device.html
type BatteryUpdate struct {
	p core.Provider

	// TimeToEmpty represented in number of seconds
	TimeToEmpty int64

	// TimeToFull represented in number of seconds
	TimeToFull int64

	// Percentage is the current percentage
	Percentage float64

	// State according to the following table:
	//   0: Unknown
	//   1: Charging
	//   2: Discharging
	//   3: Empty
	//   4: Fully charged
	//   5: Pending charge
	//   6: Pending discharge
	State uint
}

func (b *BatteryUpdate) Source() core.Provider {
	return b.p
}

func (b *BatteryUpdate) Text() string {
	charging := ""
	if b.State == 1 {
		charging = "[+]"
	}
	return fmt.Sprintf("%s%.0f%%", charging, b.Percentage)
}

const (
	dbusPropertiesInterface   = "org.freedesktop.DBus.Properties"
	dbusPropertiesGetAll      = "org.freedesktop.DBus.Properties.GetAll"
	dbusPropertiesChanged     = "PropertiesChanged"
	dbusUPowerInterface       = "org.freedesktop.UPower"
	dbusUPowerDeviceInterface = "org.freedesktop.UPower.Device"
	dbusBatteryObjectPath     = "/org/freedesktop/UPower/devices/%s"
)

type Battery struct {
	status *BatteryUpdate
	Device string `yaml:"device"`
}

func (b *Battery) Run(ch chan<- core.Update, event <-chan core.Event) {
	// open system dbus connection
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		ch <- &core.ErrorUpdate{P: b}
		return
	}
	defer conn.Close()

	objPath := dbus.ObjectPath(fmt.Sprintf(dbusBatteryObjectPath, b.Device))

	// get device initial state
	obj := conn.Object(dbusUPowerInterface, objPath)
	call := obj.Call(dbusPropertiesGetAll, 0, dbusUPowerDeviceInterface)
	if call.Err != nil {
		ch <- &core.ErrorUpdate{P: b}
		return
	}
	b.status = &BatteryUpdate{p: b}
	upd := b.extract(call.Body, 0)
	if upd == nil {
		ch <- &core.ErrorUpdate{P: b}
		return
	}
	ch <- upd
	b.status = upd

	// add signal listener for battery dbus object
	err = conn.AddMatchSignal(
		dbus.WithMatchInterface(dbusPropertiesInterface),
		dbus.WithMatchMember(dbusPropertiesChanged),
		dbus.WithMatchObjectPath(objPath),
	)
	if err != nil {
		ch <- &core.ErrorUpdate{P: b}
		return
	}

	// register signals channel with dbus
	signals := make(chan *dbus.Signal)
	conn.Signal(signals)

	// listen for signals
	for {
		signal := <-signals
		u := b.extract(signal.Body, 1)
		if u != nil {
			ch <- u
			b.status = u
		}
	}
}

func (b *Battery) extract(body []interface{}, idx int) *BatteryUpdate {
	if body == nil || len(body) < idx {
		return nil
	}
	props, ok := body[idx].(map[string]dbus.Variant)
	if !ok {
		return nil
	}
	result := b.unmarshal(props)
	result.p = b
	return result
}

// unmarshal the dbus properties.
// note: not every signal contains every field, so we have to be careful
// not to overwrite previous values with zero-values
func (b *Battery) unmarshal(props map[string]dbus.Variant) *BatteryUpdate {
	upd := b.status // copy fields from previous update
	if val, ok := props["TimeToEmpty"]; ok {
		val.Store(&upd.TimeToEmpty)
	}
	if val, ok := props["TimeToFull"]; ok {
		val.Store(&upd.TimeToFull)
	}
	if val, ok := props["Percentage"]; ok {
		val.Store(&upd.Percentage)
	}
	if val, ok := props["State"]; ok {
		val.Store(&upd.State)
	}
	return upd
}
