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
	Format          string `yaml:"date_format"`
	ShortFormat     string `yaml:"date_format_short"`
	AlternateFormat string `yaml:"date_format_alternate"`
	showAlternate   bool
}

func (c *Clock) Run(ch chan<- *core.Update, click <-chan *core.Click) {
	c.initDefaults()
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			c.sendUpdate(ch)
		case <-click:
			c.showAlternate = !c.showAlternate
			c.sendUpdate(ch)
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

func (c *Clock) sendUpdate(ch chan<- *core.Update) {
	var format string
	if c.showAlternate {
		format = c.AlternateFormat
	} else {
		format = c.Format
	}
	ch <- &core.Update{
		Source:      c,
		Status:      time.Now().Format(format),
		StatusShort: time.Now().Format(c.ShortFormat),
	}
}
