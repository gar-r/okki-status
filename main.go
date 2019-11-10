package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"bitbucket.org/dargzero/okki-status/core"
	"bitbucket.org/dargzero/okki-status/output"
)

var debug bool

var sink output.Sink
var refresh string
var cache map[core.Module]string

func main() {
	parseFlags()
	initialize()
	if refresh != "" {
		err := sendRefreshRequest(refresh)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	invalidateAll()
	go handleModuleRefresh()
	if debug {
		time.Sleep(1 * time.Minute)
	} else {
		startServer()
	}
}

func handleModuleRefresh() {
	events := make(chan core.Module, 100)
	for _, module := range config {
		module.Schedule(events)
	}
	for module := range events {
		invalidate(module)
	}
}

func invalidate(module core.Module) {
	cache[module] = module.Info()
	updateBar()
}

func invalidateAll() {
	for _, module := range config {
		cache[module] = module.Info()
	}
	updateBar()
}

func updateBar() {
	status := strings.Builder{}
	for _, entry := range config {
		_, present := cache[entry]
		if !present {
			cache[entry] = entry.Info()
		}
		status.WriteString(cache[entry])
	}
	sink.Accept(status.String())
}

func initialize() {
	if debug {
		sink = &output.StdOut{}
	} else {
		sink = &output.XRoot{}
	}
	cache = make(map[core.Module]string, len(config))
}

func parseFlags() {
	flag.BoolVar(&debug, "debug", false, "print to stdout instead of xroot")
	flag.StringVar(&refresh, "refresh", "", "refresh a single module with the give name")
	flag.Parse()
}
