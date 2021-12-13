package module

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Clock(t *testing.T) {

	timeFn = func() time.Time {
		return time.UnixMilli(0)
	}

	t.Run("no format provided", func(t *testing.T) {
		clock := &Clock{}
		s := clock.Status()
		assert.Equal(t, "Thu Jan  1 01:00:00 1970", s)
	})

	t.Run("explicit format", func(t *testing.T) {
		clock := &Clock{
			Format: "2006-01-02 3:04PM",
		}
		s := clock.Status()
		assert.Equal(t, "1970-01-01 1:00AM", s)
	})

}
