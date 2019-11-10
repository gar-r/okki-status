package output

import (
	"fmt"
	"time"
)

// StdOut is a sink that sends information to STDOUT
type StdOut struct {
	d debouncer
}

// Accept accepts the status string
func (s *StdOut) Accept(status string) {
	s.d.debounce(1*time.Second, func() {
		fmt.Println(status)
	})
}
