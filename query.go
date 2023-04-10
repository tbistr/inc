package inc

import (
	"strings"
	"unicode/utf8"
)

func (e *Engine) AddQuery(r rune) {
	e.query = append(e.query, r)

	for _, c := range e.cands {
		lPos := lastOr(c.memo.pos, 0)
		lLen := lastOr(c.memo.len, 0)
		surplus := c.Text[lPos+lLen:]
		if c.memo.matched {
			found := strings.IndexRune(surplus, r)
			if found == -1 {
				c.memo.matched = false
				continue
			}

			// head    surplus
			// "123" + "四五六"
			// if addQuery('四') ->
			// pos = lPos + lLen + found = 2 + 1 + 0 = 3
			// len = RuneLen('四') = 3
			c.memo.pos = append(c.memo.pos, lPos+lLen+uint(found))
			c.memo.len = append(c.memo.len, uint(utf8.RuneLen(r)))
		}
	}
}

func (e *Engine) RmQuery() {
	e.query = rmLast(e.query)

	for _, c := range e.cands {
		if c.memo.matched {
			c.memo.pos = rmLast(c.memo.pos)
			c.memo.len = rmLast(c.memo.len)
		}
		c.memo.matched = len(c.memo.pos) == len(e.query)
	}
}

func (e *Engine) DelQuery() {
	if len(e.query) == 0 {
		return
	}
	e.query = []rune{}

	for _, c := range e.cands {
		c.memo.matched = true
		c.memo.pos = []uint{}
		c.memo.len = []uint{}
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
