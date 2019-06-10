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
	return strings.Trim(volumeFn(), "\n")
}

var volumeFn = func() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	if err != nil {
		return "?"
	}
	return string(out)
}
