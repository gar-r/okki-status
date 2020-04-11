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
	count := countUpdates(string(out))
	return strconv.Itoa(count)
}

func countUpdates(s string) (count int) {
	for _, line := range strings.Split(s, "\n") {
		if len(strings.TrimSpace(line)) > 0 {
			count++
		}
	}
	return
}
