package inc

import "unicode/utf8"

// memo is a memoization for efficiency of transition of query add or remove.
type memo struct {
	matched bool
	// pos is the indexes of the **next** of first rune of each matched rune.
	//
	// e.g. if {query: "abc", target: "aaabbbccc"} then pos is [0, 3, 6]
	// Each index is used to get substring for next search like `target[pos[-1]+len[-1]:]]`.
	pos []uint
	// len is length of target[pos[i]].
	// It is used to get substring for next search like `target[pos[i]+len[i]]`.
	len []uint
	// NOTE: Should be []uint8? but can't calc pos + len.
}

// initMemo initializes memo for each candidate.
func (e *Engine) initMemo() {
	for i := range e.Cands {
		m, p, l := matchWithMemo(e.query, e.Cands[i].Text)
		e.Cands[i].memo = &memo{
			matched: m,
			pos:     p,
			len:     l,
		}
	}
}

// FoundRune represents a found rune info in the target.
//
// Pos is the index of the first rune of the found rune.
// Len is the length of the found rune.
//
// `target[found[i]:]` is the rest of the target(including the found rune).
// `target[found[i].Pos:found[i].Pos+found[i].Len]` is the one found rune.
type FoundRune struct {
	Pos uint
	Len uint
}

// FoundRunes returns the found runes info in the target.
//
// If the candidate is not matched, `len(FoundRunes()) < len(QueryRunes())`.
func (c Candidate) FoundRunes() []FoundRune {
	res := make([]FoundRune, 0, len(c.memo.pos))

	for i := range c.memo.pos {
		res = append(res, FoundRune{
			Pos: c.memo.pos[i],
			Len: c.memo.len[i],
		})
	}
	return res
}

// matchWithMemo does naive incremental search and returns memo for it.
func matchWithMemo(query []rune, target string) (matched bool, poss []uint, lens []uint) {
	// Ok to specify cap = len(query), but increase memory usage.
	if len(query) == 0 {
		return true, poss, lens
	}

	byteI := uint(0)
	cursor := 0
	for _, r := range target {
		if r == query[cursor] {
			cursor++
			poss = append(poss, byteI)
			lens = append(lens, uint(utf8.RuneLen(r)))
		}
		if cursor == len(query) {
			return true, poss, lens
		}
		byteI += uint(utf8.RuneLen(r))
	}

	return false, poss, lens
}
