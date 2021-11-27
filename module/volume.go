package module

import (
	"os/exec"
	"strings"

	"hu.okki.okki-status/core"
)

// Volume provides system volume related information
type Volume struct {
}

// Status returns the primary device volume in percentage, or the muted string
func (v *Volume) Status() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	volStr := strings.Trim(string(out), "\n")
	if failed(volStr, err) {
		return core.StatusUnknown
	}
	return volStr
}

func failed(volStr string, err error) bool {
	return err != nil && volStr != "muted" && volStr != "0%"
}
