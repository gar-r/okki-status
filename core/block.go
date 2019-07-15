package core

import (
	"strings"
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

type Block struct {
	Gap            Gap
	BlockOrder     BlockOrder
	IconProvider   IconProvider
	StatusProvider StatusProvider
}

func (b *Block) String() string {
	status := b.StatusProvider.GetStatus()
	icon := b.IconProvider.GetIcon(status)
	if b.BlockOrder == IconFirst {
		return b.Gap.Format(icon, status)
	}
	return b.Gap.Format(status, icon)
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
