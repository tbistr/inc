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
	// e.g. if {query: "abc", target: "aaabbbccc"} then starts is [1, 4, 7].
	// Each index is used to get substring for next search like `target[starts[len(starts)-1]:]]`.
	starts []uint
}

func (e *Engine) initMemo() {
	for i := range e.cands {
		m, is := matchWithMemo(e.query, e.cands[i].Text)
		e.cands[i].memo = &memo{
			matched: m,
			starts:  is,
		}
	}
}

func matchWithMemo(query []rune, target string) (bool, []uint) {
	starts := make([]uint, 0, len(query))
	if len(query) == 0 {
		return true, starts
	}

	byteI := uint(0)
	cursor := 0
	for _, r := range target {
		byteI += uint(utf8.RuneLen(r))
		if r == query[cursor] {
			cursor++
			starts = append(starts, byteI)
		}
		if cursor == len(query) {
			return true, starts
		}
	}

	return false, starts
}
