package core

import (
	"strings"
	"time"
)

type StatusProvider interface {
	GetStatus() string
}

type IconProvider interface {
	GetIcon(status string) string
}

type BlockOrder int

const (
	IconFirst = iota
	TextFirst
)

type Module struct {
	Gap            Gap
	BlockOrder     BlockOrder
	IconProvider   IconProvider
	StatusProvider StatusProvider
	Frequency      time.Duration
}

type Message struct {
	M Module
	S string
}

func (m Module) Schedule(ch chan Message) {
	ticker := time.NewTicker(m.Frequency)
	go func() {
		for range ticker.C {
			msg := Message{
				M: m,
				S: m.Info(),
			}
			ch <- msg
		}
	}()
}

func (m *Module) Info() string {
	status := m.StatusProvider.GetStatus()
	icon := m.IconProvider.GetIcon(status)
	if m.BlockOrder == IconFirst {
		return m.Gap.Format(icon, status)
	}
	return m.Gap.Format(status, icon)
}

type Gap struct {
	Before, After string
}

func (g *Gap) Format(values ...string) string {
	sb := &strings.Builder{}
	sb.WriteString(g.Before)
	for _, value := range values {
		sb.WriteString(value)
	}
	sb.WriteString(g.After)
	return sb.String()
}
