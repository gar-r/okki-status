package core

import (
	"testing"
)

func Test_NewBar(t *testing.T) {
	modules := make([]Module, 0)
	bar := NewBar(modules)
	if bar == nil {
		t.Errorf("expected non-nill value")
	}
}

func Test_Render(t *testing.T) {
	modules := []Module{
		makeTestModule("mod1", "status1"),
		makeTestModule("mod2", "status2"),
	}
	bar := NewBar(modules)

	var actual string
	bar.Render(func(s string) {
		actual = s
	})

	expected := "|status1||status2|"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func Test_Invalidate(t *testing.T) {
	module := makeTestModule("mod", "status")
	bar := NewBar([]Module{module})

	bar.Invalidate(module)
}
