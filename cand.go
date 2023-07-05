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

// KeyRune represents a found rune info in the target.
//
// `target[found.Pos:]` is the rest of the target(including the found rune).
// `target[found.Pos:found.Pos+found.Len]` is the one found rune.
type KeyRune struct {
	// Pos is the indexes of the **next** of first rune of each matched rune.
	//
	// e.g. if {query: "abc", target: "aaabbbccc"} then Pos is [0, 3, 6]
	// Each index is used to get substring for next search like `target[Pos[-1]+len[-1]:]]`.
	Pos uint
	// Len is length of target[Pos].
	// It is used to get substring for next search like `target[pos+Len]`.
	Len uint
	// NOTE: Should be []uint8? but can't calc pos + len.
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
