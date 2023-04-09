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
	indexes []uint
}

func (e *Engine) initMemo() {
	for _, cand := range e.cands {
		m, is := matchWithMemo(e.query, cand.Text)
		cand.memo = &memo{
			matched: m,
			indexes: is,
		}
	}
}

func matchWithMemo(query []rune, target string) (bool, []uint) {
	indexes := make([]uint, 0, len(query))
	if len(query) == 0 {
		return true, indexes
	}

	byteI := uint(0)
	cursor := 0
	for _, r := range target {
		byteI += uint(utf8.RuneLen(r))
		if r == query[cursor] {
			cursor++
			indexes = append(indexes, byteI)
		}
		if cursor == len(query) {
			return true, indexes
		}
	}

	return false, indexes
}
