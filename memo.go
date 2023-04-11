package inc

import (
	"unicode/utf8"
)

// memo is a memoization for efficiency of transition of query add or remove.
type memo struct {
	matched bool
	founds  []FoundRune
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

// initMemo initializes memo for each candidate.
func (e *Engine) initMemo() {
	for i := range e.Cands {
		e.Cands[i].memo = matchWithMemo(e.query, e.Cands[i].Text)
	}
}

// FoundRunes returns the found runes info in the target.
//
// If the candidate is not matched, `len(FoundRunes()) < len(QueryRunes())`.
func (c Candidate) FoundRunes() []FoundRune {
	// Shallow copy is acceptable.
	return append([]FoundRune{}, c.memo.founds...)
}

// matchWithMemo does naive incremental search and returns memo for it.
func matchWithMemo(query []rune, target string) *memo {
	// Ok to specify cap = len(query), but increase memory usage.
	founds := make([]FoundRune, 0)
	if len(query) == 0 {
		return &memo{true, founds}
	}

	byteI := uint(0)
	cursor := 0
	for _, r := range target {
		if r == query[cursor] {
			cursor++
			founds = append(founds, FoundRune{
				byteI, uint(utf8.RuneLen(r)),
			})
		}
		if cursor == len(query) {
			return &memo{true, founds}
		}
		byteI += uint(utf8.RuneLen(r))
	}

	return &memo{false, founds}
}
