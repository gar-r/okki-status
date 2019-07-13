package modules

import (
	"math"
	"os/exec"
	"strconv"
	"strings"
)

// Brightness provides display brightness related status using brillo
type Brightness struct {
	Margin
	Icon
}

// Status provides the display brightness status
func (b *Brightness) Status() string {
	val := b.value()
	return b.Margin.Format(b.Icon.Format(val))
}

func (b *Brightness) value() string {
	out, err := exec.Command("brillo", "-G").Output()
	if err != nil {
		return "?"
	}
	valStr := strings.Trim(string(out), "\n")
	f, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return "?"
	}
	return strconv.Itoa(int(math.Round(f)))
}
