package module

import (
	"fmt"
	"os/exec"
	"regexp"

	"hu.okki.okki-status/core"
)

var ssidRe = regexp.MustCompile(`SSID:\s+(.*)`)
var signalRe = regexp.MustCompile(`signal:\s+(.*)`)

// WiFi provides wireless network information for the given device
type WiFi struct {
	Device string
}

// Status returns the connected WiFi SSID name and the signal strength
func (w *WiFi) Status() string {
	info := w.getInfo()
	if info == nil {
		return core.StatusError
	}
	ssid := w.ssid(info)
	if ssid != "" {
		signal := w.signal(info)
		if signal != "" {
			return fmt.Sprintf("%s (%s)", ssid, signal)
		}
		return ssid
	}
	return core.StatusUnknown
}

func (w *WiFi) ssid(info []byte) string {
	return w.findFirst(info, ssidRe)
}

func (w *WiFi) signal(info []byte) string {
	return w.findFirst(info, signalRe)
}

func (w *WiFi) findFirst(info []byte, re *regexp.Regexp) string {
	if match := re.FindSubmatch(info); len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func (w *WiFi) getInfo() []byte {
	out, err := exec.Command("iw", "dev", w.Device, "link").Output()
	if err != nil {
		return nil
	}
	return out
}
