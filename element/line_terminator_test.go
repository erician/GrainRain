package element

import "testing"

func TestIsLineTerminatorWithSimpleCases(t *testing.T) {
	cases := []struct {
		in 		rune;
		want 	bool;
	}{
		{'\n', true},
		{'\r', true},
		{'a', false},
	}
	for _, c := range cases {
		got := IsWhiteSpace(c.in)
		if got != c.want {
			t.Errorf("IsWhiteSpace(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}