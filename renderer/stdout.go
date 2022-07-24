package renderer

import (
	"fmt"

	"okki-status/core"
)

// StdOut renders the bar using the standard output
type StdOut struct {
	Separator  string // rendered between blocks
	Terminator string // rendered after the last block
}

func (s *StdOut) Render(bar core.Bar) {
	for i, block := range bar {
		stdout(block.Status())
		if i+1 == len(bar) {
			stdout(s.Terminator)
		} else {
			stdout(s.Separator)
		}
	}
}

// allow override stdout fn for testing
var stdout = func(s string) {
	fmt.Print(s)
}
