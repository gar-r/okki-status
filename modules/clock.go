package modules

import "time"

// Clock provides date/time status in a given Layout
type Clock struct {
	Layout string
}

// Status returns the date/time in the format specified by Layout
func (c *Clock) Status() string {
	if c.Layout == "" {
		c.Layout = time.ANSIC
	}
	return time.Now().Format(c.Layout)
}
