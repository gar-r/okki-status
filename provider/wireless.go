package provider

import (
	"errors"
	"fmt"
	"okki-status/core"
	"time"

	"github.com/mdlayher/wifi"
)

type WirelessUpdate struct {
	p              core.Provider
	SSID           string
	SignalStrength int
}

func (w *WirelessUpdate) Source() core.Provider {
	return w.p
}

func (w *WirelessUpdate) Text() string {
	return fmt.Sprintf("%s (%ddBm)", w.SSID, w.SignalStrength)
}

type Wireless struct {
	InterfaceName string `yaml:"interface_name"`
	Refresh       int    `yaml:"refresh"`
}

func (w *Wireless) Run(ch chan<- core.Update, events <-chan core.Event) {
	ch <- w.getWifiInfo()

	if w.Refresh == 0 {
		w.Refresh = 5000
	}
	t := time.NewTicker(time.Duration(w.Refresh) * time.Millisecond)
	for {
		select {
		case <-t.C:
			ch <- w.getWifiInfo()
		case e := <-events:
			if _, ok := e.(*core.Refresh); ok {
				ch <- w.getWifiInfo()
			}
		}
	}
}

func (w *Wireless) getWifiInfo() core.Update {
	client, err := wifi.New()
	if err != nil {
		return &core.ErrorUpdate{P: w}
	}
	defer client.Close()
	iface, err := w.findInterface(client)
	if err != nil {
		return &core.ErrorUpdate{P: w}
	}
	bss, err := client.BSS(iface)
	if err != nil {
		return &core.ErrorUpdate{P: w}
	}
	sinfo, err := client.StationInfo(iface)
	if err != nil || len(sinfo) == 0 {
		return &core.ErrorUpdate{P: w}
	}
	return &WirelessUpdate{
		p:              w,
		SSID:           bss.SSID,
		SignalStrength: sinfo[0].Signal,
	}
}

func (w *Wireless) findInterface(client *wifi.Client) (*wifi.Interface, error) {
	interfaces, err := client.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range interfaces {
		if i.Name == w.InterfaceName {
			return i, nil
		}
	}
	return nil, errors.New("wireless interface not found: " + w.InterfaceName)
}
