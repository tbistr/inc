package inc

type candidate struct {
	ptr   *any
	text  string
	match bool
}

type Engine struct {
	cs    []candidate
	query string
}

func New(query string, candidates []struct {
	ptr  *any
	text string
}) *Engine {
	cs := make([]candidate, len(candidates))
	for i, c := range candidates {
		cs[i] = candidate{
			ptr:   c.ptr,
			text:  c.text,
			match: Match(query, c.text),
		}
	}

	return &Engine{
		cs:    cs,
		query: query,
	}
}

func Match(query string, body string) bool {
	if len(query) == 0 {
		return true
	}

	cursor := 0
	for _, c := range body {
		if c == rune(query[cursor]) {
			cursor++
		}
		if cursor == len(query) {
			return true
		}
	}

	return false
}

func (e *Engine) Matched() ([]int, []string) {
	is := make([]int, 0, len(e.cs))
	ss := make([]string, 0, len(e.cs))

	for i := 0; i < len(e.cs); i++ {
		if e.cs[i].match {
			is = append(is, i)
			ss = append(ss, e.cs[i].text)
		}
	}
	return is, ss
}

func (e *Engine) MatchedPtr() []*any {
	res := make([]*any, 0, len(e.cs))

	for i := 0; i < len(e.cs); i++ {
		res = append(res, e.cs[i].ptr)
	}
	return res
}

func (e *Engine) AddQuery(c rune) {
	e.query += string(c)
	for _, c := range e.cs {
		c.match = Match(e.query, c.text)
	}
}

func (e *Engine) DelQuery() {
	if len(e.query) == 0 {
		return
	}

	e.query = e.query[:len(e.query)-1]
	for _, c := range e.cs {
		c.match = Match(e.query, c.text)
	}
}
