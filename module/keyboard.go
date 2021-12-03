package module

import (
	"os/exec"
	"regexp"

	"hu.okki.okki-status/core"
)

var layoutRe = regexp.MustCompile(`.*layout:\s*([a-z]*)`)

// Keyboard provides keyboard layout information
type Keyboard struct {
}

// Status returns the keyboard layout string
func (l *Keyboard) Status() string {
	info, err := exec.Command("setxkbmap", "-query").Output()
	if err != nil {
		return core.StatusError
	}
	return layout(info)
}

func layout(info []byte) string {
	if match := layoutRe.FindSubmatch(info); len(match) >= 2 {
		return string(match[1])
	}
	return core.StatusError
}
