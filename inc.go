package inc

type Candidate struct {
	Ptr  any
	Text string
	memo *memo
}

type Engine struct {
	cands []Candidate
	query []rune
}

func New(query string, cands []Candidate) *Engine {
	e := &Engine{
		cands: cands,
		query: []rune(query),
	}
	e.initMemo()
	return e
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
	res := make([]int, 0, len(e.cands))

	for i := range e.cands {
		if e.cands[i].memo.matched {
			res = append(res, i)
		}
	}
	return res
}

func (e *Engine) MatchedString() []string {
	res := make([]string, 0, len(e.cands))

	for _, c := range e.cands {
		if c.memo.matched {
			res = append(res, c.Text)
		}
	}
	return res
}

func (e *Engine) MatchedPtr() []any {
	res := make([]any, 0, len(e.cands))

	for _, c := range e.cands {
		res = append(res, c.Ptr)
	}
	return res
}
