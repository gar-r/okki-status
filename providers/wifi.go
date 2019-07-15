package providers

import (
	"fmt"
	"os/exec"
	"regexp"
)

type Wifi struct {
	Device string
}

func (w *Wifi) GetStatus() string {
	info := w.getInfo()
	if info == nil {
		return ":("
	}
	ssid := w.ssid(info)
	if ssid != "" {
		signal := w.signal(info)
		if signal != "" {
			return fmt.Sprintf("%s (%s)", ssid, signal)
		}
		return ssid
	}
	return "?"
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

func (w *Wifi) getInfo() []byte {
	out, err := exec.Command("iw", "dev", w.Device, "link").Output()
	if err != nil {
		return nil
	}
	return out
}
