package output

import (
	"fmt"
)

// StdOut is a sink that sends information to STDOUT
type StdOut struct {
	d debouncer
}

// Accept accepts the status string
func (s *StdOut) Accept(status string) {
	s.d.invoke(func() {
		fmt.Println(status)
	})
}
