package output

import (
	"sync"
	"time"
)

type State int

const (
	IDLE State = iota
	INITIAL
	DEBOUNCE
)

const interval = 100 * time.Millisecond

type debouncer struct {
	state  State
	timer  *time.Timer
	cached func()
	lock   sync.Mutex
}

func (d *debouncer) invoke(f func()) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.cached = f
	switch d.state {
	case IDLE:
		d.state = INITIAL
		d.startTimer()
		d.flush()
	case INITIAL:
		d.state = DEBOUNCE
		d.timer.Reset(interval)
	case DEBOUNCE:
		d.timer.Reset(interval)
		d.cached = f
	}
}

func (d *debouncer) timerElapsed() {
	d.lock.Lock()
	defer d.lock.Unlock()
	switch d.state {
	case INITIAL:
		d.state = IDLE
	case DEBOUNCE:
		d.state = IDLE
		d.flush()
	}
}

func (d *debouncer) startTimer() {
	if d.timer == nil {
		d.timer = time.NewTimer(interval)
		go func() {
			for range d.timer.C {
				d.timerElapsed()
			}
		}()
	} else {
		d.timer.Reset(interval)
	}
}

func (d *debouncer) flush() {
	go d.cached()
}
