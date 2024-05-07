package kebab

import (
	"testing"
)

type kebabTest struct {
	s, expected string
}

func TestToKebabCase(t *testing.T) {
	items := []kebabTest{
		{"What is programming?", "what-is-programming?"},
		{"already-in-kebab-case", "already-in-kebab-case"},
		{"Can make a URL", "can-make-a-url"},
	}

	for _, it := range items {
		got := ToKebabCase(it.s)
		want := it.expected

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}
}
