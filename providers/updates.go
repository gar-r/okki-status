package providers

import (
	"os/exec"
	"strings"
)

// Updates provides the number of available updates
type Updates struct {
}

// GetStatus returns the number of updates as string
func (u *Updates) GetStatus() string {
	out, err := exec.Command("yay", "-Pn").Output()
	if err != nil {
		return "?"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) < 1 {
		return "?"
	}
	return lines[0]
}
