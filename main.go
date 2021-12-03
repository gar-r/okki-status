package main

import (
	"time"

	"hu.okki.okki-status/config"
	"hu.okki.okki-status/refresh"
)

func main() {
	startRenderLoop()
	refresh.Listen()
}

func startRenderLoop() {
	t := time.NewTicker(config.Interval)
	go func() {
		for range t.C {
			config.R.Render(config.B)
		}
	}()
}
