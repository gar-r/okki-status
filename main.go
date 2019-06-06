package main

import (
	"os"
	"strings"
	"time"

	"bitbucket.org/dargzero/smart-status/sinks"
)

func main() {
	sink := initSink()
	for {
		status := getStatus()
		sink.Accept(status)
		time.Sleep(1 * time.Second)
	}
}

func getStatus() string {
	status := strings.Builder{}
	for _, module := range activeModules {
		status.WriteString(module.Status())
		status.WriteString(separator)
	}
	return status.String()
}

func initSink() sinks.Sink {
	var sink sinks.Sink
	if isDebugMode() {
		sink = &sinks.StdOut{}
	} else {
		sink = &sinks.Xroot{}
	}
	return sink
}

func isDebugMode() bool {
	return len(os.Args) >= 2 && os.Args[1] == "debug"
}
