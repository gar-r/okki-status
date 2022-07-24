package module

import (
	"math"
	"os/exec"
	"strconv"
	"strings"

	"okki-status/core"
)

// Brightness provides screen brightness information
type Brightness struct {
}

// Status returns the display brightness in percentage
func (b *Brightness) Status() string {
	out, err := exec.Command("brillo", "-G").Output()
	if err != nil {
		return core.StatusError
	}
	valStr := strings.Trim(string(out), "\n")
	f, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return core.StatusError
	}
	return strconv.Itoa(int(math.Round(f)))
}
