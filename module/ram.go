package module

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"hu.okki.okki-status/core"
)

// RAM provides system memory related information
type RAM struct {
}

// Status returns the used system memory percentage
func (r *RAM) Status() string {
	raw, err := ramCmd.Exec()
	if err != nil {
		return core.StatusError
	}
	var re = regexp.MustCompile(`Mem:\s+(\d+)\s+(\d+)`)
	if match := re.FindSubmatch(raw); len(match) >= 3 {
		// below errors can be ignored, since regex already matched decimal
		total, _ := strconv.ParseFloat(string(match[1]), 64)
		used, _ := strconv.ParseFloat(string(match[2]), 64)
		percentUsed := math.Round(used / total * 100)
		return fmt.Sprintf("%.0f%%", percentUsed)
	}
	return core.StatusError
}

var ramCmd core.Command = &core.OsCommand{
	Name: "free",
	Args: []string{"--bytes"},
}
