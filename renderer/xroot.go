package renderer

import (
	"log"
	"os/exec"
	"strings"

	"okki-status/core"
)

// XRoot renders the bar using the xsetroot Xorg utility
type XRoot struct {
	Separator  string // rendered between blocks
	Terminator string // rendered after the last block
}

func (x *XRoot) Render(bar core.Bar) {
	sb := strings.Builder{}
	for i, block := range bar {
		sb.WriteString(block.Status())
		if i+1 == len(bar) {
			sb.WriteString(x.Terminator)
		} else {
			sb.WriteString(x.Separator)
		}
	}
	x.set(sb.String())
}

func (x *XRoot) set(s string) {
	err := exec.Command("xsetroot", "-name", s).Run()
	if err != nil {
		log.Println(err)
	}
}
