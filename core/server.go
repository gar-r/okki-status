package core

import (
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const (
	okkiStatusInterface     = "hu.okki.garric.OkkiStatus"
	okkiStatusObj           = "/hu/okki/garric/OkkiStatus"
	introspectableInterface = "org.freedesktop.DBus.Introspectable"
)

const intro = `
<node>
  <interface name="hu.okki.garric.OkkiStatus">
    <method name="Refresh">
      <arg direction="in" type="s" />
    </method>
  </interface>` + introspect.IntrospectDataString + `</node>`

type dbusServer struct {
	bar *Bar
}

func (d *dbusServer) Refresh(moduleName string) *dbus.Error {
	for _, m := range d.bar.Modules {
		if m.Name == moduleName {
			m.events <- &Refresh{
				Name: m.Name,
			}
		}
	}
	return nil
}

func (b *Bar) startDbusServer() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		b.errors <- err
		return
	}
	defer conn.Close()

	d := &dbusServer{bar: b}
	conn.Export(d, okkiStatusObj, okkiStatusInterface)
	conn.Export(introspect.Introspectable(intro), okkiStatusObj, introspectableInterface)

	res, err := conn.RequestName(okkiStatusInterface, dbus.NameFlagDoNotQueue)
	if err != nil {
		b.errors <- err
		return
	}
	if res != dbus.RequestNameReplyPrimaryOwner {
		b.errors <- errors.New("dbus name already taken")
		return
	}
	select {}
}
