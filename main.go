package main

import (
	"flag"

	"hu.okki.okki-status/core"
)

var debug bool
var outputFn func(string)

var bar = core.NewBar(modules)
var events = make(chan core.Module, 100)

func main() {
	parseFlags()
	initSink()
	go waitForEvents()
	schedule()
	listenForExternal()
}

func parseFlags() {
	flag.BoolVar(&debug, "debug", false, "print to stdout instead of xroot")
	flag.Parse()
}

func initSink() {
	if debug {
		outputFn = stdout
	} else {
		outputFn = xroot
	}
	bar.Render(outputFn)
}

func schedule() {
	for _, module := range modules {
		module.Schedule(events)
	}
}

func waitForEvents() {
	for module := range events {
		bar.Invalidate(module)
		bar.Render(outputFn)
	}
}
