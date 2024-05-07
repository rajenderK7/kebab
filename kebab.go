package kebab

import "strings"

const (
	hyphen = '-'
)

var sb = &strings.Builder{}

func ToKebabCase(s string) string {
	sb.Reset() // Empty the buffer.
	for _, r := range s {
		if isSpace(r) {
			sb.WriteRune(hyphen)
		} else if !isAlpha(r) {
			// Special characters...
			sb.WriteRune(r)
		} else {
			if !isLower(r) {
				r = toLower(r)
			}
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
