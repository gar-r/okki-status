package main

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"bitbucket.org/dargzero/smart-status/sinks"
)

var debugMode bool
var command string
var args []string

func main() {
	readArgs()
	sink := initSink()
	if command != "" {
		execCommand()
		updateStatus(sink)
		return
	}
	for {
		updateStatus(sink)
		time.Sleep(1 * time.Minute)
	}
}

func updateStatus(sink sinks.Sink) {
	status := getStatus()
	sink.Accept(status)
}

func execCommand() {
	err := exec.Command(command, args...).Run()
	if err != nil {
		os.Exit(1)
	}
}

func getStatus() string {
	status := strings.Builder{}
	for _, module := range config {
		status.WriteString(module.Status())
	}
	return status.String()
}

func initSink() sinks.Sink {
	var sink sinks.Sink
	if debugMode {
		sink = &sinks.StdOut{}
	} else {
		sink = &sinks.Xroot{}
	}
	return sink
}

func readArgs() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		if i == 1 {
			if arg == "-d" || arg == "--debug" {
				debugMode = true
			} else {
				command = arg
			}
		} else {
			args = append(args, arg)
		}
	}
}
