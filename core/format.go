package core

import "strings"

type BlockOrder int

const (
	IconFirst = iota
	TextFirst
)

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
