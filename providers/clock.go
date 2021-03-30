package providers

import "time"

// Clock provides date/time information in a given layout
type Clock struct {
	Layout string
}

// GetStatus returns the date/time string in the given format
func (c *Clock) GetStatus() string {
	if c.Layout == "" {
		c.Layout = time.ANSIC
	}
	return time.Now().Format(c.Layout)
}
