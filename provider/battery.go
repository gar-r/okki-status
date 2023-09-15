package provider

import (
	"okki-status/core"
	"time"
)

type Battery struct {
	Device string `yaml:"device"`
}

func (b *Battery) Run(c chan<- *core.Event) {
	for {
		c <- &core.Event{
			Source: b,
			Status: time.Now().String(),
		}
		time.Sleep(time.Minute)
	}
}
