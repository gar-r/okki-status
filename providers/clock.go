package providers

import "time"

// Clock provides date/time status in a given Layout
type Clock struct {
	Layout string
}

// GetData returns the provided date/time status
func (c *Clock) GetData() string {
	if c.Layout == "" {
		c.Layout = time.ANSIC
	}
	return time.Now().Format(c.Layout)
}
