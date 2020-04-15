package core

import (
	"time"
)

// Module represents a single module on the status bar
type Module struct {
	Name       string
	Gap        Gap
	BlockOrder BlockOrder
	Icon       IconProvider
	Status     StatusProvider
	Refresh    time.Duration
	Delay      time.Duration
}

// Info retrieves the and formats the status bar information
func (m *Module) Info() string {
	status := m.Status.GetStatus()
	icon := m.Icon.GetIcon(status)
	if m.BlockOrder == IconFirst {
		return m.Gap.Format(icon, status)
	}
	return m.Gap.Format(status, icon)
}

// Schedule creates a ticker to refesh the module periodically
func (m Module) Schedule(ch chan Module) {
	ticker := time.NewTicker(m.Refresh)
	go func() {
		for range ticker.C {
			ch <- m
		}
	}()
}
