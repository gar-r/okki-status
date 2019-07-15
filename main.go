package main

import (
	"bitbucket.org/dargzero/smart-status/core"
	"bitbucket.org/dargzero/smart-status/output"
	"flag"
	"strings"
)

var debug bool
var single bool

var sink output.Sink
var cache map[core.Module]string

func main() {
	parseFlags()
	initialize()
	emit()
	if single {
		return
	}
	listen()
}

func listen() {
	ch := make(chan core.Message)
	for _, entry := range config {
		entry.Schedule(ch)
	}
	for msg := range ch {
		cache[msg.M] = msg.S
		emit()
	}
}

func emit() {
	status := strings.Builder{}
	for _, entry := range config {
		s, present := cache[entry]
		if !present {
			info := entry.Info()
			cache[entry] = info
			status.WriteString(info)
		} else {
			status.WriteString(s)
		}

	}
	sink.Accept(status.String())
}

func initialize() {
	if debug {
		sink = &output.StdOut{}
	} else {
		sink = &output.Xroot{}
	}
	cache = make(map[core.Module]string, len(config))
}

func parseFlags() {
	flag.BoolVar(&debug, "debug", false, "print to stdout instead of xroot")
	flag.BoolVar(&single, "single", false, "refresh all modules and exit")
	flag.Parse()
}
