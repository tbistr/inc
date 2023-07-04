package inc

// Candidate is a candidate for incremental search.
//
// Text is a string to be searched.
// If you want to Engine to return a pointer to some object as a search result, you can use the Ptr field.
type Candidate struct {
	ptr      any
	text     []rune
	matched  bool
	keyRunes []FoundRune
}

// FoundRune represents a found rune info in the target.
//
// `target[found.Pos:]` is the rest of the target(including the found rune).
// `target[found.Pos:found.Pos+found.Len]` is the one found rune.
type FoundRune struct {
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
		ptr:     ptr,
		text:    text,
		matched: true,
	}
}

func (c *Candidate) String() string {
	if c == nil {
		return ""
	}
	return string(c.text)
}

// Matched returns whether the candidate is matched.
func (c *Candidate) Matched() bool {
	if c == nil {
		return false
	}
	return c.matched
}

func (c *Candidate) GetKeyRunes() []FoundRune {
	if c == nil {
		return nil
	}
	return c.keyRunes
}

// InnerCandidate is a inner representation of Candidate.
//
// It is equivalent to *Candidate, but it has a method to get/set private fields.
// Assume that it is used from Algorithm.
//
// TODO: Make InnerCandidate private.
// This type can be used from user code, but it should be private from user.
// But it must can be used from `/algorithm/foo`.
// Define it in something like `/internal/interface` package?
// However, it occurs a import cycle.
// How can I do this?
type InnerCandidate struct {
	*Candidate
}

func (c InnerCandidate) SetMatched(matched bool) {
	if c.Candidate == nil {
		return
	}
	c.matched = matched
}

func (c InnerCandidate) SetKeyRunes(keyRunes []FoundRune) {
	if c.Candidate == nil {
		return
	}
	c.keyRunes = keyRunes
}
