package core

import "testing"

func Test_Gap_Format(t *testing.T) {
	gap := Gap{
		Before: "|",
		After:  "|",
	}

	actual := gap.Format("foo", " bar", " baz")
	expected := "|foo bar baz|"

	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func Test_NtoI(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"0", 0},
		{"45", 45},
		{"-17", -17},
		{"75%", 75},
		{"silly", -1},
		{"13+", -1},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := NtoI(tc.input)
			if actual != tc.output {
				t.Errorf("expected %v, got %v", tc.output, actual)
			}
		})
	}

}
