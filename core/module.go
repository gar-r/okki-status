package core

import (
	"time"
)

type Module struct {
	Name       string
	Gap        Gap
	BlockOrder BlockOrder
	Icon       IconProvider
	Status     StatusProvider
	Refresh    time.Duration
}

func (m *Module) Info() string {
	status := m.Status.GetStatus()
	icon := m.Icon.GetIcon(status)
	if m.BlockOrder == IconFirst {
		return m.Gap.Format(icon, status)
	}
	return m.Gap.Format(status, icon)
}

func (m Module) Schedule(ch chan Module) {
	ticker := time.NewTicker(m.Refresh)
	go func() {
		for range ticker.C {
			ch <- m
		}
	}()
}
