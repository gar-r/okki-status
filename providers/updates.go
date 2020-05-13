package providers

import (
	"os/exec"
	"strconv"
	"strings"
)

// Updates provides the number of available updates
type Updates struct {
	// The command to execute: it should return each update on a separate line
	Command string

	// Arguments for the command
	Args []string
}

// GetStatus returns the number of updates as string
func (u *Updates) GetStatus() string {
	out, err := exec.Command(u.Command, u.Args...).Output()
	if err != nil {
		return "?"
	}
	count := countNonEmptyLines(string(out))
	return strconv.Itoa(count)
}

func countNonEmptyLines(s string) int {
	lines := strings.Split(string(s), "\n")
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	return count
}
