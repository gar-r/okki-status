package core

import (
	"strconv"
	"strings"
)

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
func (g *Gap) Format(str ...string) string {
	sb := &strings.Builder{}
	sb.WriteString(g.Before)
	for _, value := range str {
		sb.WriteString(value)
	}
	sb.WriteString(g.After)
	return sb.String()
}

var DefaultGap = Gap{
	Before: "   ",
	After:  "   ",
}

// NtoI converts a number-string to int - the number string can also
// be a percentage. For example "45%"" will become the int value of 45.
func NtoI(value string) int {
	n, err := strconv.Atoi(strings.Replace(value, "%", "", 1))
	if err != nil {
		return -1
	}
	return n
}
