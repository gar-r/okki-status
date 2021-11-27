package module

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"strconv"

	"hu.okki.okki-status/core"
)

// RAM provides system memory related information
type RAM struct {
}

// Status returns the used system memory percentage
func (r *RAM) Status() string {
	raw, err := r.getInfo()
	if err != nil {
		return core.StatusError
	}
	var re = regexp.MustCompile(`Mem:\s+(\d+)\s+(\d+)`)
	if match := re.FindSubmatch(raw); len(match) >= 3 {
		total, err := strconv.ParseFloat(string(match[1]), 64)
		if err != nil {
			return core.StatusError
		}
		used, err := strconv.ParseFloat(string(match[2]), 64)
		if err != nil {
			return core.StatusError
		}
		percentUsed := math.Round(used / total * 100)
		return fmt.Sprintf("%.0f%%", percentUsed)
	}
	return core.StatusError
}

func (r *RAM) getInfo() ([]byte, error) {
	out, err := exec.Command("free", "--bytes").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
