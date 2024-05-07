package kebab

import (
	"testing"
)

type kebabTest struct {
	s, expected string
}

func TestToKebabCase(t *testing.T) {
	items := []kebabTest{
		{"already-in-kebab-case", "already-in-kebab-case"},
		{"What is programming?", "what-is-programming?"},
		{"ALL UPPER CASE", "all-upper-case"},
		{"This include7s 69", "this-include7s-69"},
		{"$peci@l ch@r@cter$", "$peci@l-ch@r@cter$"},
	}

	for _, it := range items {
		got := ToKebabCase(it.s)
		want := it.expected

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}
}
