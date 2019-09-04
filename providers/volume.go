package providers

import (
	"os/exec"
	"strings"
)

// Volume provides system volume related information
type Volume struct {
}

// GetStatus returns the primary device volume in percentage, or the muted string
func (v *Volume) GetStatus() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	volStr := strings.Trim(string(out), "\n")
	if failed(volStr, err) {
		return "?"
	}
	return volStr
}

func failed(volStr string, err error) bool {
	return err != nil && volStr != "muted" && volStr != "0%"
}
