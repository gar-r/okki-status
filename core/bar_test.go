package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Block(t *testing.T) {

	mod := &tester{T: "foo"}

	t.Run("simple block", func(t *testing.T) {
		b := &Block{Module: mod}
		s := b.Status()
		assert.Equal(t, "foo", s)
	})

	t.Run("block with prefix and postfix", func(t *testing.T) {
		b := &Block{
			Module:  mod,
			Prefix:  "===",
			Postfix: "---",
		}
		s := b.Status()
		assert.Equal(t, "===foo---", s)
	})

}

type tester struct {
	T string
}

func (t *tester) Status() string {
	return t.T
}
