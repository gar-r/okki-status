package modules

import (
	"time"
)

// Clock provides date/time status in a given Layout
type Clock struct {
	Margin
	Icon
	Layout string
}

// Status returns the date/time in the format specified by Layout
func (c *Clock) Status() string {
	if c.Layout == "" {
		c.Layout = time.ANSIC
	}
	clock := time.Now().Format(c.Layout)
	return c.Margin.Format(c.Icon.Format(clock))
}
