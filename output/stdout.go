package output

import "fmt"

// StdOut is a sink that sends information to STDOUT
type StdOut struct {
}

// Accept accepts the status string
func (*StdOut) Accept(status string) {
	fmt.Println(status)
}
