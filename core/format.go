package core

import "strings"

// BlockOrder determines the order in which the text and icon is printed
type BlockOrder int

const (
	// IconFirst renders the icon before the text
	IconFirst = iota

	// TextFirst renders the text before the icon
	TextFirst
)

// Gap contains the gap characters added around a block
type Gap struct {
	Before, After string
}

// Format prints the block formatted with gaps
func (g *Gap) Format(values ...string) string {
	sb := &strings.Builder{}
	sb.WriteString(g.Before)
	for _, value := range values {
		sb.WriteString(value)
	}
	sb.WriteString(g.After)
	return sb.String()
}
