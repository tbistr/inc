package inc

// Candidate is a candidate for incremental search.
//
// Text is a string to be searched.
// If you want to Engine to return a pointer to some object as a search result, you can use the Ptr field.
type Candidate struct {
	Ptr      any
	Text     []rune
	Matched  bool
	KeyRunes []KeyRune
}

// KeyRune represents a important rune in a candidate.Text.
// It is intended to be used to highlight matched runes.
//
// Use like this:
//
//	target[Pos:]
//	target[Pos:Pos+Len]
//
// Former is the substring from the key rune to the end of the target.
// Latter is one key runes with proper length.
//
// If multibyte (non-ASCII) characters are not used, you may use only Pos.
type KeyRune struct {
	// Pos is the index of the first rune of the runes which are matched.
	Pos uint
	// Len is length of the rune starts with target[Pos].
	Len uint
}

// NewCandidate Makes a new Candidate.
//
// text is a string to be searched.
//
// ptr is a arbitrary pointer.
// If you want to Engine to return a pointer to some object as a search result, you can use this field.
func NewCandidate(text []rune, ptr any) *Candidate {
	return &Candidate{
		Ptr:     ptr,
		Text:    text,
		Matched: true,
	}
}

func (c *Candidate) String() string {
	if c == nil {
		return ""
	}
	return string(c.Text)
}
