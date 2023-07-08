package inc

// Candidate is a candidate for incremental search.
type Candidate struct {
	// Ptr is a arbitrary pointer.
	// If you want to Engine to return a pointer to some object as a search result, you can use this field.
	Ptr any
	// Text is a string to be searched.
	Text []rune
	// Matched is true if the candidate is matched.
	Matched bool
	// KeyRunes is a list of key runes.
	// It is intended to be used to highlight matched runes.
	//
	// Even if Matched is true, len(KeyRunes) != len(query) may be true.
	// It depends on the algorithm you use.
	// (Some algorithms may not use KeyRunes at all.)
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
func NewCandidate(text []rune, ptr any) *Candidate {
	return &Candidate{
		Ptr:     ptr,
		Text:    text,
		Matched: true,
	}
}

func (c Candidate) String() string {
	return string(c.Text)
}
