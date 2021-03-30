package core

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_StaticIcon(t *testing.T) {
	i := &StaticIcon{
		Icon: "X",
	}

	str := i.GetIcon("")

	if str != "X" {
		t.Errorf("expected %v, got %v", "X", str)
	}
}

func Test_ThresholdIcon(t *testing.T) {
	i := &ThresholdIcon{
		Thresholds: []Threshold{
			{90, "A"},
			{50, "B"},
			{10, "C"},
		},
		StatusConverterFn: func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		},
	}

	tests := []struct {
		input, output string
	}{
		{"0", ""},
		{"9", ""},
		{"10", "C"},
		{"11", "C"},
		{"49", "C"},
		{"50", "B"},
		{"51", "B"},
		{"89", "B"},
		{"90", "A"},
		{"91", "A"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v => %v", tc.input, tc.output), func(t *testing.T) {
			str := i.GetIcon(tc.input)
			if str != tc.output {
				t.Errorf("expected %v, got %v", tc.output, str)
			}
		})
	}

}
