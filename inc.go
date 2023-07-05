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

type Algorithm interface {
	// AppendCands appends candidates to the engine.
	AppendCands([]*Candidate)

	// GetQuery returns the current query.
	GetQuery() []rune
	// AddQuery adds a rune to the query.
	AddQuery(rune)
	// RmQuery removes a rune from the query.
	RmQuery()
	// DelQuery deletes (clears) the query.
	DelQuery()
}

// Engine is a engine for incremental search.
// Cands is a list of candidates.
type Engine struct {
	// Candidate should not be pointer because it can be touched from user in a process.
	cands []*Candidate
	Algorithm
}

// New returns a new Engine.
func New(query string, cands []Candidate) *Engine {
	e := &Engine{Algorithm: &DefaultAlgo{}}
	e.AppendCands(cands)
	for _, r := range query {
		e.AddQuery(r)
	}
	return e
}

func (e *Engine) AppendCands(cands []Candidate) {
	pCands := make([]*Candidate, len(cands))
	for i, c := range cands {
		pCands[i] = &c
	}
	e.cands = append(e.cands, pCands...)
}

// Matched returns matched candidates.
func (e *Engine) Matched() []Candidate {
	matched := make([]Candidate, 0, len(e.cands))
	for _, c := range e.cands {
		if c.Matched {
			matched = append(matched, *c)
		}
	}
	return matched
}
