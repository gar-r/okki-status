package providers

import (
	"os/exec"
	"strconv"
	"strings"
)

// Updates provides the number of available updates
type Updates struct {
}

// GetStatus returns the number of updates as string
func (u *Updates) GetStatus() string {
	out, err := exec.Command("yay", "-Qu").Output()
	if err != nil {
		return "?"
	}
	n := len(strings.Split(string(out), "\n"))
	return strconv.Itoa(n)
}
