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

// Candidate is a candidate for incremental search.
//
// Text is a string to be searched.
// If you want to Engine to return a pointer to some object as a search result, you can use the Ptr field.
type Candidate struct {
	Ptr  any
	Text string
	memo *memo
}

// Engine is a engine for incremental search.
// Cands is a list of candidates.
type Engine struct {
	Cands []Candidate
	query []rune
}

// New returns a new Engine.
func New(query string, cands []Candidate) *Engine {
	e := &Engine{
		Cands: cands,
		query: []rune(query),
	}
	e.initMemo()
	return e
}

// Query returns the current query string.
func (e *Engine) Query() string {
	return string(e.query)
}

// Matched returns whether the candidate is matched.
func (c *Candidate) Matched() bool {
	return c.memo.matched
}

// MatchedIndex returns the index of the matched candidates.
func (e *Engine) MatchedIndex() []int {
	res := make([]int, 0, len(e.Cands))

	for i := range e.Cands {
		if e.Cands[i].memo.matched {
			res = append(res, i)
		}
	}
	return res
}

// MatchedString returns the text of the matched candidates.
func (e *Engine) MatchedString() []string {
	res := make([]string, 0, len(e.Cands))

	for _, c := range e.Cands {
		if c.memo.matched {
			res = append(res, c.Text)
		}
	}
	return res
}

// MatchedPtr returns the pointer of the matched candidates.
//
// Ptr is Cands.Ptr.
func (e *Engine) MatchedPtr() []any {
	res := make([]any, 0, len(e.Cands))

	for _, c := range e.Cands {
		if c.memo.matched {
			res = append(res, c.Ptr)
		}
	}
	return res
}
