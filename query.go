package inc

import (
	"unicode/utf8"
)

// AddQueryString adds a string to the query.
func (e *Engine) AddQueryString(s string) {
	for _, r := range s {
		e.AddQuery(r)
	}
}

// AddQuery adds a rune to the query.
func (e *Engine) AddQuery(r rune) {
	e.query = append(e.query, r)

	for _, c := range e.Cands {
		if c.memo.matched {
			last := lastOr(c.memo.founds, FoundRune{0, 0})
			surplus := c.Text[last.Pos+last.Len:]
			found := e.index(surplus, r)
			if found == -1 {
				c.memo.matched = false
				continue
			}

			// head    surplus
			// "123" + "四五六"
			// if addQuery('四') ->
			// Pos = lPos + lLen + found = 2 + 1 + 0 = 3
			// Len = RuneLen('四') = 3
			c.memo.founds = append(c.memo.founds, FoundRune{
				last.Pos + last.Len + uint(found),
				uint(utf8.RuneLen(r)),
			})
		}
	}
}

// RmNQuery removes the last n runes from the query.
func (e *Engine) RmNQuery(n int) {
	if n <= 0 {
		return
	}
	for i := 0; i < n; i++ {
		e.RmQuery()
	}
}

// RmQuery removes the last rune from the query.
func (e *Engine) RmQuery() {
	e.query = rmLast(e.query)

	for _, c := range e.Cands {
		if c.memo.matched {
			c.memo.founds = rmLast(c.memo.founds)
		}
		c.memo.matched = len(c.memo.founds) == len(e.query)
	}
}

// DelQuery removes all runes from the query.
// All candidates will be matched.
func (e *Engine) DelQuery() {
	if len(e.query) == 0 {
		return
	}
	e.query = []rune{}

	for _, c := range e.Cands {
		c.memo = &memo{true, []FoundRune{}}
	}
}

func lastOr[T any](ts []T, defaultV T) T {
	if len(ts) == 0 {
		return defaultV
	}
	return ts[len(ts)-1]
}

func rmLast[T any](ts []T) []T {
	if len(ts) == 0 {
		return ts
	}
	return ts[:len(ts)-1]
}
