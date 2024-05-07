package kebab

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isAlpha(r rune) bool {
	return isLower(r) || isUpper(r)
}

func isSpace(r rune) bool {
	return r == ' '
}

func toLower(r rune) rune {
	if isLower(r) {
		return r
	}
	// The ASCII code of upper case 'A' is 65.
	// We add the offset to lower case 'a' to get the
	// lower case version of the given rune.
	return 'a' + (r - 65)
}
