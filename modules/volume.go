package modules

import (
	"os/exec"
	"strings"
)

// Volume provides volume related status using pamixer
type Volume struct {
	Margin
	Icon
}

// Status returns the volume status
func (v *Volume) Status() string {
	vol := v.getValue()
	return v.Margin.Format(v.Icon.Format(vol))
}

func (v *Volume) getValue() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	volStr := strings.Trim(string(out), "\n")
	if err != nil && volStr != "muted" {
		return "?"
	}
	return volStr
}
