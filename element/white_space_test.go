package element

import "testing"

func TestIsWhiteSpaceWithSimpleCases(t *testing.T) {
	cases := []struct {
		in 		rune;
		want 	bool;
	}{
		{'\n', true},
		{'a', false},
	}
	for _, c := range cases {
		got := IsWhiteSpace(c.in)
		if got != c.want {
			t.Errorf("IsWhiteSpace(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestIsWhiteSpaceWithAString(t *testing.T) {
	input := "hello你好\t\n "
	wants := []bool{false, false, false, false, false, 
		false, false, true, true, true}
	// rune is traversed by code point and the index are always right,
	// so we should use rune.
	// for more details: https://www.jianshu.com/p/f397a35360da
	for i, r := range []rune(input) {
		got := IsWhiteSpace(r)
		if got != wants[i] {
			t.Errorf("IsWhiteSpace(%d) == %t, want %t", r, got, wants[i])
		}
	}
}