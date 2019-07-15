package providers

import (
	"os/exec"
	"strings"
)

type Volume struct {
}

func (v *Volume) GetStatus() string {
	out, err := exec.Command("pamixer", "--get-volume-human").Output()
	volStr := strings.Trim(string(out), "\n")
	if err != nil && volStr != "muted" {
		return "?"
	}
	return volStr
}
