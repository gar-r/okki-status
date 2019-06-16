package modules

import (
	"os/exec"
	"strings"
)

// Volume provides volume related status using pamixer
type Volume struct {
}

// Status returns the volume status
func (v *Volume) Status() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	volStr := strings.Trim(string(out), "\n")
	if err != nil && volStr != "muted" {
		return "?"
	}
	return volStr
}
