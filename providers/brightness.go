package providers

import (
	"math"
	"os/exec"
	"strconv"
	"strings"
)

type Brightness struct {
}

func (b *Brightness) GetStatus() string {
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
