package main

import (
	"time"

	"okki-status/config"
	"okki-status/refresh"
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
