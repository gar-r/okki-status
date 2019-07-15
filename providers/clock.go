package providers

import "time"

type Clock struct {
	Layout string
}

func (c *Clock) GetStatus() string {
	if c.Layout == "" {
		c.Layout = time.ANSIC
	}
	return time.Now().Format(c.Layout)
}
