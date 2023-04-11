package inc

type Candidate struct {
	Ptr  any
	Text string
	memo *memo
}

type Engine struct {
	Cands []Candidate
	query []rune
}

func New(query string, cands []Candidate) *Engine {
	e := &Engine{
		Cands: cands,
		query: []rune(query),
	}
	e.initMemo()
	return e
}

func (e *Engine) Query() string {
	return string(e.query)
}

func Match(query string, body string) bool {
	return matchWithRune([]rune(query), body)
}

func matchWithRune(query []rune, body string) bool {
	if len(query) == 0 {
		return true
	}

	cursor := 0
	for _, c := range body {
		if c == query[cursor] {
			cursor++
		}
		if cursor == len(query) {
			return true
		}
	}

	return false
}

func (e *Engine) MatchedIndex() []int {
	res := make([]int, 0, len(e.Cands))

	for i := range e.Cands {
		if e.Cands[i].memo.matched {
			res = append(res, i)
		}
	}
	return res
}

func (e *Engine) MatchedString() []string {
	res := make([]string, 0, len(e.Cands))

	for _, c := range e.Cands {
		if c.memo.matched {
			res = append(res, c.Text)
		}
	}
	return res
}

func (e *Engine) MatchedPtr() []any {
	res := make([]any, 0, len(e.Cands))

	for _, c := range e.Cands {
		if c.memo.matched {
			res = append(res, c.Ptr)
		}
	}
	return res
}
