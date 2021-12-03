package module

import "time"

// Clock provides date/time information in a given layout
type Clock struct {
	Format string // optional format string for datetime
}

// Status returns the date/time string in the given format
func (c *Clock) Status() string {
	if c.Format == "" {
		c.Format = time.ANSIC
	}
	return time.Now().Format(c.Format)
}
