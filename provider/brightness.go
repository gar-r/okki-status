package provider

import (
	"fmt"
	"okki-status/core"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Brightness struct {
	Refresh int `yaml:"refresh"`
}

func (b *Brightness) Run(ch chan<- core.Update, event <-chan core.Event) {
	ch <- b.getUpdate()

	if b.Refresh == 0 {
		b.Refresh = 60000
	}

	ticker := time.NewTicker(time.Duration(b.Refresh) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			ch <- b.getUpdate()
		case e := <-event:
			if _, ok := e.(*core.Refresh); ok {
				ch <- b.getUpdate()
			}
		}
	}
}

func (b *Brightness) getUpdate() core.Update {
	out, err := exec.Command("brillo", "-G").Output()
	if err != nil {
		return &core.ErrorUpdate{P: b}
	}
	valStr := strings.TrimSpace(string(out))
	f, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return &core.ErrorUpdate{P: b}
	}
	return &core.SimpleUpdate{
		P: b,
		T: fmt.Sprintf("%.0f%%", f),
	}
}
