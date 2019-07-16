package strutil

import "strings"

func IsUpper(character rune) bool {
	return strings.ContainsRune(UpperCharacters, character)
}
