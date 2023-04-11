package inc

import "strings"

// index is custom search function.
type index func(s string, r rune) int

// initIndex initializes the index function.
func (e *Engine) initIndex() {
	e.index = strings.IndexRune
	if e.option.ignoreCase {
		e.index = indexIgnoreCase
	}
}

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
