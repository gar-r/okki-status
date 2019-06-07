package modules

import (
	"fmt"
	"os/exec"
	"regexp"
)

// Wifi provides Wifi related information
type Wifi struct {
	Device string
}

// Status returns network name and signal strength
func (w *Wifi) Status() string {
	info := wifiFn(w.Device)
	if info == nil {
		return ""
	}
	ssid := w.ssid(info)
	signal := w.signal(info)
	return fmt.Sprintf("%s (%s)", ssid, signal)
}

func (w *Wifi) ssid(info []byte) string {
	var re = regexp.MustCompile(`SSID:\s+(.*)`)
	return w.findFirst(info, re)
}

func (w *Wifi) signal(info []byte) string {
	var re = regexp.MustCompile(`signal:\s+(.*)`)
	return w.findFirst(info, re)
}

func (w *Wifi) findFirst(info []byte, re *regexp.Regexp) string {
	if match := re.FindSubmatch(info); len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

var wifiFn = func(device string) []byte {
	out, err := exec.Command("iw", "dev", device, "link").Output()
	if err != nil {
		return nil
	}
	return out
}
