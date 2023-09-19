package provider

import (
	"okki-status/core"
	"time"
)

const (
	DefaultFormat      = "2006-01-02 15:04:05"
	DefaultShortFormat = "15:04"
)

type Clock struct {
	Format          string `yaml:"clock_format"`
	ShortFormat     string `yaml:"clock_format_short"`
	AlternateFormat string `yaml:"clock_format_alternate"`
	showAlternate   bool
}

func (c *Clock) Run(ch chan<- core.Update, event <-chan core.Event) {
	c.initDefaults()
	c.sendUpdate(ch) // send an initial update
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			c.sendUpdate(ch)
		case e := <-event:
			if _, ok := e.(*core.Click); ok {
				c.showAlternate = !c.showAlternate
				c.sendUpdate(ch)
			}
		}
	}
}

func (c *Clock) initDefaults() {
	if c.Format == "" {
		c.Format = DefaultFormat
	}
	if c.ShortFormat == "" {
		c.ShortFormat = DefaultShortFormat
	}
}

func (c *Clock) sendUpdate(ch chan<- core.Update) {
	var format string
	if c.showAlternate {
		format = c.AlternateFormat
	} else {
		format = c.Format
	}
	ch <- &core.SimpleUpdate{
		P: c,
		T: time.Now().Format(format),
	}
}
