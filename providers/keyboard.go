package providers

import (
	"os/exec"
	"regexp"
)

var layoutRe = regexp.MustCompile(`.*layout:\s*([a-z]*)`)

type Layout struct {
}

func (l *Layout) GetStatus() string {
	info, err := exec.Command("setxkbmap", "-query").Output()
	if err != nil {
		return "?"
	}
	return layout(info)
}

func layout(info []byte) string {
	if match := layoutRe.FindSubmatch(info); len(match) >= 2 {
		return string(match[1])
	}
	return "?"
}
