package inc

// Match does naive incremental search.
func Match(query string, body string) bool {
	if len(query) == 0 {
		return true
	}

	queryRunes := []rune(query)
	cursor := 0
	for _, c := range body {
		if c == queryRunes[cursor] {
			cursor++
		}
		if cursor == len(query) {
			return true
		}
	}

	return false
}

// Engine is a engine for incremental search.
type Engine struct {
	cands []*Candidate
	Algorithm
}

// New returns a new Engine.
//
// It uses the default algorithm which is naive incremental search.
func New(query string, cands []Candidate) *Engine {
	return NewWithAlgo(query, cands, &defaultAlgo{})
}

// NewWithAlgo returns a new Engine with a custom algorithm.
func NewWithAlgo(query string, cands []Candidate, algo Algorithm) *Engine {
	e := &Engine{Algorithm: algo}
	e.AppendCands(cands)
	for _, r := range query {
		e.AddQuery(r)
	}
	return e
}

// AppendCands appends candidates to the engine.
//
// It receives candidates as values and converts them to pointers.
// So, you can't modify the candidates after passing them to AppendCands.
func (e *Engine) AppendCands(cands []Candidate) {
	pCands := make([]*Candidate, len(cands))
	for i, c := range cands {
		p := c // Loop var c is overwritten in each iteration.
		pCands[i] = &p
	}

	e.Algorithm.AppendCands(pCands)
	e.cands = append(e.cands, pCands...)
}

// Matched returns matched candidates.
//
// It returns candidates as values, so you can't modify internal states of the engine.
func (e *Engine) Matched() []Candidate {
	matched := make([]Candidate, 0, len(e.cands))
	for _, c := range e.cands {
		if c.Matched {
			matched = append(matched, *c)
		}
	}
	return matched
}
