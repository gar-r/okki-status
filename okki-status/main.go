package main

import (
	"flag"
	"time"

	"hu.okki.okki-status/core"
	"hu.okki.okki-status/output"
)

var debug bool
var sink output.Sink

var bar = core.NewBar(modules)
var events = make(chan core.Module, 100)

func main() {
	parseFlags()
	initSink()
	go waitForEvents()
	schedule()
	initServer()
}

func parseFlags() {
	flag.BoolVar(&debug, "debug", false, "print to stdout instead of xroot")
	flag.Parse()
}

func initSink() {
	if debug {
		sink = &output.StdOut{}
	} else {
		sink = &output.XRoot{}
	}
	bar.Render(sink)
}

func schedule() {
	for _, module := range modules {
		module.Schedule(events)
	}
}

func waitForEvents() {
	for module := range events {
		bar.Invalidate(module)
		bar.Render(sink)
	}
}

func initServer() {
	if !debug {
		listenForExternal()
	} else {
		time.Sleep(5 * time.Minute)
	}
}
