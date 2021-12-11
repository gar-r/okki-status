package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Text(t *testing.T) {
	m := &Text{T: "foo"}
	s := m.Status()

	assert.Equal(t, "foo", s)
}
