package modules

import (
	"testing"
)

func TestVolume_Status(t *testing.T) {
	volumeFn = func() string {
		return "test\n"
	}

	v := &Volume{}

	expected := "test"
	actual := v.Status()

	if actual != "test" {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}
