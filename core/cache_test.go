package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CachingModule(t *testing.T) {

	const never = 100 * time.Hour

	t.Run("status refreshed on creation", func(t *testing.T) {
		m := &Tester{Mock: "foo"}
		cm := NewCachingModule(m, never)
		s := cm.Status()
		assert.Equal(t, m.Mock, s)
	})

	t.Run("invalidate status", func(t *testing.T) {
		m := &Tester{Mock: "foo"}
		cm := NewCachingModule(m, never)

		m.Mock = "bar"
		cm.Invalidate()
		s := cm.Status()

		assert.Equal(t, m.Mock, s)
	})

	t.Run("status cached", func(t *testing.T) {
		m := &Tester{Mock: "foo"}
		cm := NewCachingModule(m, never)

		s1 := cm.Status()
		m.Mock = "bar"
		s2 := cm.Status()

		assert.Equal(t, s1, s2)

	})

	t.Run("cache expires", func(t *testing.T) {
		m := &Tester{Mock: "foo"}
		expiry := 100 * time.Millisecond
		cm := NewCachingModule(m, expiry)

		s1 := cm.Status()
		m.Mock = "bar"
		time.Sleep(2 * expiry)
		s2 := cm.Status()

		assert.Equal(t, "foo", s1)
		assert.Equal(t, "bar", s2)
	})

}

type Tester struct {
	Mock string
}

func (t *Tester) Status() string {
	return t.Mock
}
