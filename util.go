package inc

func Strs2Cands(ts []string) []Candidate {
	cands := make([]Candidate, 0, len(ts))
	for i, s := range ts {
		cands = append(cands, Candidate{
			Ptr:  i,
			Text: s,
		})
	}
	return cands
}
