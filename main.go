package main

import (
	"bitbucket.org/dargzero/smart-status/core"
	"bitbucket.org/dargzero/smart-status/output"
	"flag"
	"log"
	"strings"
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
	startServer()
}

func handleModuleRefresh() {
	ch := make(chan core.Module)
	for _, module := range config {
		module.Schedule(ch)
	}
	for module := range ch {
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
		sink = &output.Xroot{}
	}
	cache = make(map[core.Module]string, len(config))
}

func parseFlags() {
	flag.BoolVar(&debug, "debug", false, "print to stdout instead of xroot")
	flag.StringVar(&refresh, "refresh", "", "refresh a single module with the give name")
	flag.Parse()
}
