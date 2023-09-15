package provider

import (
	"okki-status/core"
	"time"
)

type Clock struct{}

func (c *Clock) Run(ch chan<- *core.Update, click <-chan *core.Click) {
	format1 := "2006-01-02"
	format2 := "15:04:05"
	f := format1
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			ch <- &core.Update{
				Source: c,
				Status: time.Now().Format(f),
			}
		case <-click:
			if f == format1 {
				f = format2
			} else {
				f = format1
			}
			ch <- &core.Update{
				Source: c,
				Status: time.Now().Format(f),
			}
		}
	}
}
