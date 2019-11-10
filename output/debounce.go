package output

import (
	"sync"
	"time"
)

type debouncer struct {
	m     sync.Mutex
	timer *time.Timer
}

func (d *debouncer) debounce(interval time.Duration, f func()) {
	d.m.Lock()
	defer d.m.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(interval, f)
}
