package main

import (
	"fmt"
	"os"
	"strings"

	"bitbucket.org/dargzero/smart-status/sinks"
)

func main() {
	config := NewConfig()
	sink := initSink()
	status := strings.Builder{}
	for _, entry := range config.Entries {
		status.WriteString(getFormattedData(entry))
	}
	sink.Accept(status.String())
}

func initSink() sinks.Sink {
	var sink sinks.Sink
	if os.Args[1] == "debug" {
		sink = &sinks.StdOut{}
	} else {
		sink = &sinks.Xroot{}
	}
	return sink
}

func getFormattedData(entry Entry) string {
	data := entry.provider.GetData()
	if entry.format != "" {
		data = fmt.Sprintf(entry.format, data)
	}
	return data
}
