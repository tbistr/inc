package inc

type candidate struct {
	ptr   *any
	text  string
	match bool
}

type Window struct {
	cs    []candidate
	query string
}

func New(query string, candidates []struct {
	ptr  *any
	text string
}) *Window {
	cs := make([]candidate, len(candidates))
	for i, c := range candidates {
		cs[i] = candidate{
			ptr:   c.ptr,
			text:  c.text,
			match: Match(query, c.text),
		}
	}

	return &Window{
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

func (w *Window) Matched() ([]int, []string) {
	is := make([]int, 0, len(w.cs))
	ss := make([]string, 0, len(w.cs))

	for i := 0; i < len(w.cs); i++ {
		if w.cs[i].match {
			is = append(is, i)
			ss = append(ss, w.cs[i].text)
		}
	}
	return is, ss
}

func (w *Window) MatchedPtr() []*any {
	res := make([]*any, 0, len(w.cs))

	for i := 0; i < len(w.cs); i++ {
		res = append(res, w.cs[i].ptr)
	}
	return res
}

func (w *Window) AddQuery(c rune) {
	w.query += string(c)
	for _, c := range w.cs {
		c.match = Match(w.query, c.text)
	}
}

func (w *Window) DelQuery() {
	if len(w.query) == 0 {
		return
	}

	w.query = w.query[:len(w.query)-1]
	for _, c := range w.cs {
		c.match = Match(w.query, c.text)
	}
}
