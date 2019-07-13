package modules

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"strconv"
)

// RAM module provides information on RAM usage
type RAM struct {
	Margin
	Icon
}

// Status returns the memory status
func (r *RAM) Status() string {
	v := r.getValue()
	return r.Margin.Format(r.Icon.Format(v))
}

func (r *RAM) getValue() string {
	raw, err := r.getInfo()
	const errValue = ":("
	if err != nil {
		return errValue
	}
	var re = regexp.MustCompile(`Mem:\s+(\d+)\s+(\d+)`)
	if match := re.FindSubmatch(raw); len(match) >= 3 {
		total, err := strconv.ParseFloat(string(match[1]), 64)
		if err != nil {
			return errValue
		}
		used, err := strconv.ParseFloat(string(match[2]), 64)
		if err != nil {
			return errValue
		}
		percentUsed := math.Round(used / total * 100)
		return fmt.Sprintf("%.0f%%", percentUsed)
	}
	return errValue
}

func (r *RAM) getInfo() ([]byte, error) {
	out, err := exec.Command("free", "--bytes").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
