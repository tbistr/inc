package inc

import "unicode/utf8"

type memo struct {
	// NOTE: Should really define memo as struct?
	// matched can be defined in Candidate.
	// Then indexes can be defined as memo(field).
	matched bool
	// whole   string
	// surplus string
	// surplus is calculated by target[indexes[len(indexes)-1]:]

	// starts is the indexes of the **next** of first rune of each matched rune.
	//
	// e.g. if {query: "abc", target: "aaabbbccc"} then pos is [0, 3, 6]
	// Each index is used to get substring for next search like `target[starts[len(starts)-1]:]]`.
	pos []uint
	// len is length of target[pos[i]].
	// It is used to get substring for next search like `target[pos[i]+len[i]]`.
	len []uint
	// NOTE: Should be []uint8? but can't calc pos + len.
}

func (e *Engine) initMemo() {
	for i := range e.cands {
		m, p, l := matchWithMemo(e.query, e.cands[i].Text)
		e.cands[i].memo = &memo{
			matched: m,
			pos:     p,
			len:     l,
		}
	}
}

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
