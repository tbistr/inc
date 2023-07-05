package inc

// Strs2Cands converts a list of strings to a list of Candidates.
//
// Ptr of each Candidate will be set to nil for decreasing memory usage.
func Strs2Cands(ts []string) []Candidate {
	cands := make([]Candidate, 0, len(ts))
	for _, s := range ts {
		cands = append(cands, *NewCandidate([]rune(s), nil))
	}
	return cands
}
