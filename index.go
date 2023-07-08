package inc

// TODO: Enable custom index option.

import "strings"

// Index is custom search function.
type Index func(s string, r rune) int

// indexIgnoreCase is a custom search function that ignores case.
// Only ASCII characters are supported.
//
// A <-> a
func indexIgnoreCase(s string, r rune) int {
	switch {
	case 'a' <= r && r <= 'z':
		return strings.IndexFunc(s, func(ir rune) bool {
			return ir == r || ir == r+('A'-'a')
		})
	case 'A' <= r && r <= 'Z':
		return strings.IndexFunc(s, func(ir rune) bool {
			return ir == r || ir == r+('a'-'A')
		})
	default:
		return strings.IndexRune(s, r)
	}
}
