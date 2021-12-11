package module

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"hu.okki.okki-status/core"
)

func Test_Ram(t *testing.T) {

	ram := &RAM{}

	t.Run("ram command error", func(t *testing.T) {
		ramCmd = &core.ErrCommand{
			Err: errors.New("test"),
		}
		s := ram.Status()
		assert.Equal(t, core.StatusError, s)
	})

	t.Run("ram parse succesful", func(t *testing.T) {
		ramCmd = &core.ConstCommand{
			Output: `      total        used        free      shared  buff/cache   available
			Mem:      8244359168  1339752448  4913147904   417157120  1991458816  6225346560
			Swap:     8489267200           0  8489267200`,
		}
		s := ram.Status()
		assert.Equal(t, "16%", s)
	})

}
