package modules

import (
	"testing"
	"time"
)

func TestClock_Status(t *testing.T) {

	t.Run("default layout", func(t *testing.T) {
		c := &Clock{}
		data := c.Status()
		if !canParse(data, c.Layout) {
			t.Errorf("unable to parse %s", data)
		}
	})

	t.Run("explicit layout", func(t *testing.T) {
		c := &Clock{Layout: time.Kitchen}
		data := c.Status()
		if !canParse(data, time.Kitchen) {
			t.Errorf("unable to parse %s", data)
		}
	})

}

func canParse(s string, layout string) bool {
	_, err := time.Parse(layout, s)
	return err == nil
}
