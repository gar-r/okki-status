package output

import (
	"sync"
	"time"
)

// State represents the current state in the state machine
type State int

const (
	// Idle State means, that the debouncer is idle
	Idle State = iota

	// Initial State means, that the debouncer is not started yet
	Initial

	// Debounce State means, that the debouncer is currently suppressing events
	Debounce
)

const interval = 1 * time.Second

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
	case Idle:
		d.state = Initial
		d.startTimer()
		d.flush()
	case Initial:
		d.state = Debounce
		d.timer.Reset(interval)
	case Debounce:
		d.timer.Reset(interval)
		d.cached = f
	}
}

func (d *debouncer) timerElapsed() {
	d.lock.Lock()
	defer d.lock.Unlock()
	switch d.state {
	case Initial:
		d.state = Idle
	case Debounce:
		d.state = Idle
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
