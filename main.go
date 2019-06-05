package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"bitbucket.org/dargzero/smart-status/sinks"
)

func main() {
	config := NewConfig()
	sink := initSink()
	for {
		status := getStatus(config)
		sink.Accept(status)
		time.Sleep(1 * time.Second)
	}
}

func getStatus(config *Config) string {
	status := strings.Builder{}
	for _, entry := range config.Entries {
		status.WriteString(getFormattedData(entry))
	}
	return status.String()
}

func getFormattedData(entry Entry) string {
	data := entry.provider.GetData()
	if entry.format != "" {
		data = fmt.Sprintf(entry.format, data)
	}
	return data
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
