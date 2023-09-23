package provider

import (
	"fmt"
	"okki-status/core"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type VolumeUpdate struct {
	p            core.Provider
	sink         string
	source       string
	SinkVolume   int
	SourceVolume int
	SinkMuted    bool
	SourceMuted  bool
}

func NewVolumeUpdate(sinkStr, sourceStr string) *VolumeUpdate {
	v := &VolumeUpdate{
		sink:   sinkStr,
		source: sourceStr,
	}
	v.SinkVolume, v.SinkMuted = parseVol(v.sink)
	if v.source != "" {
		v.SourceVolume, v.SourceMuted = parseVol(v.source)
	}
	return v
}

func parseVol(volStr string) (int, bool) {
	matches := volRe.FindAllStringSubmatch(volStr, -1)
	if len(matches) == 0 {
		return 0, false
	}
	f, err := strconv.ParseFloat(matches[0][1], 32)
	if err != nil {
		return 0, false
	}
	b := len(matches[0][2]) > 0
	return int(float32(f) * 100), b
}

func (v *VolumeUpdate) Source() core.Provider {
	return v.p
}

func (v *VolumeUpdate) Text() string {
	var sink string
	if v.SinkMuted {
		sink = fmt.Sprintf("(M)%s", v.sink)
	} else {
		sink = v.sink
	}
	if v.source == "" {
		return sink
	}
	if v.SourceMuted {
		return fmt.Sprintf("%s, (M)%s", sink, v.source)
	}
	return fmt.Sprintf("%s, %s", sink, v.source)
}

type Volume struct {
	ticker  *time.Ticker
	Refresh int  `yaml:"refresh"`
	ShowMic bool `yaml:"show_mic"`
}

func (v *Volume) Run(ch chan<- core.Update, event <-chan core.Event) {
	ch <- v.getVolumeUpdate()

	if v.Refresh == 0 {
		v.Refresh = 60000
	}
	v.ticker = time.NewTicker(time.Duration(v.Refresh) * time.Millisecond)
	for {
		select {
		case <-v.ticker.C:
			ch <- v.getVolumeUpdate()
		case e := <-event:
			if c, ok := e.(*core.Click); ok {
				handleClick(c)
			} else if _, ok := e.(*core.Refresh); ok {
				ch <- v.getVolumeUpdate()
			}
		}
	}
}

func (v *Volume) getVolumeUpdate() core.Update {
	sink, err := v.getVolume(defaultSink)
	if err != nil {
		return &core.ErrorUpdate{P: v}
	}
	source := ""
	if v.ShowMic {
		source, err = v.getVolume(defaultSource)
		if err != nil {
			return &core.ErrorUpdate{P: v}
		}
	}
	upd := NewVolumeUpdate(sink, source)
	upd.p = v
	return upd
}

func (v *Volume) getVolume(dev string) (string, error) {
	out, err := exec.Command("wpctl", "get-volume", dev).Output()
	if err != nil {
		return "", err
	}
	vol := string(out)
	return strings.TrimSpace(vol), nil
}

func handleClick(c *core.Click) {
	var dev string
	if c.Button == 1 {
		dev = defaultSink
	} else {
		dev = defaultSource
	}
	_, _ = exec.Command("wpctl", "set-mute", dev, "toggle").Output()
}

const (
	volPattern    = `^Volume:\s*(\d.\d\d)\s*(\[MUTED\])*$`
	defaultSink   = "@DEFAULT_AUDIO_SINK@"
	defaultSource = "@DEFAULT_AUDIO_SOURCE@"
)

var volRe = regexp.MustCompile(volPattern)
