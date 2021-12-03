package renderer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"hu.okki.okki-status/core"
	"hu.okki.okki-status/module"
)

func Test_StdOut(t *testing.T) {

	// mock render fn
	actual := strings.Builder{}
	stdout = func(s string) {
		actual.WriteString(s)
	}

	bar := core.Bar{
		&core.Block{Module: &module.Text{T: "foo"}},
		&core.Block{Module: &module.Text{T: "bar"}},
		&core.Block{Module: &module.Text{T: "baz"}},
	}

	t.Run("basic render", func(t *testing.T) {
		actual.Reset()
		renderer := &StdOut{}
		renderer.Render(bar)
		assert.Equal(t, "foobarbaz", actual.String())
	})

	t.Run("render with prefix", func(t *testing.T) {
		actual.Reset()
		renderer := &StdOut{Separator: "|"}
		renderer.Render(bar)
		assert.Equal(t, "foo|bar|baz", actual.String())
	})

	t.Run("basic render", func(t *testing.T) {
		actual.Reset()
		renderer := &StdOut{Terminator: "!"}
		renderer.Render(bar)
		assert.Equal(t, "foobarbaz!", actual.String())
	})

}
