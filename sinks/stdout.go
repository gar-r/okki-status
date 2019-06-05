package sinks

import "fmt"

type StdOut struct {
}

func (*StdOut) Accept(status string) {
	fmt.Printf(status)
}
