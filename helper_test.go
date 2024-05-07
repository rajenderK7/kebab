package kebab

import "testing"

type boolTest struct {
	r        rune
	expected bool
}

type charTest struct {
	r, expected rune
}

func TestIsLower(t *testing.T) {
	items := []boolTest{
		{'r', true},
		{'7', false},
		{'P', false},
		{'@', false},
	}

	for _, it := range items {
		got := isLower(it.r)
		want := it.expected

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

func TestIsUpper(t *testing.T) {
	items := []boolTest{
		{'r', false},
		{'7', false},
		{'P', true},
		{'@', false},
	}

	for _, it := range items {
		got := isUpper(it.r)
		want := it.expected

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	items := []boolTest{
		{'r', true},
		{'7', false},
		{'P', true},
		{'@', false},
	}

	for _, it := range items {
		got := isAlpha(it.r)
		want := it.expected

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

func TestIsSpace(t *testing.T) {
	items := []boolTest{
		{' ', true},
		{'7', false},
		{'P', false},
		{'@', false},
	}

	for _, it := range items {
		got := isSpace(it.r)
		want := it.expected

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	}
}

func TestToLower(t *testing.T) {
	chars := []charTest{
		{'r', 'r'},
		{'V', 'v'},
	}

	for _, ch := range chars {
		got := toLower(ch.r)
		want := ch.expected

		if got != want {
			t.Errorf("got %c, wanted %c", got, want)
		}
	}
}
