package inc

import (
	"strings"
	"unicode/utf8"
)

func (e *Engine) AddQuery(r rune) {
	e.query = append(e.query, r)

	for _, c := range e.cands {
		lastStart := lastOr(c.memo.starts, 0)
		surplus := c.Text[lastStart:]
		if c.memo.matched {
			found := strings.IndexRune(surplus, r)
			if found == -1 {
				c.memo.matched = false
				continue
			}

			// head    surplus
			// "123" + "四五六"
			// len('四') == 3
			// search '四' -> uint(len(surplus)+found+utf8.RuneLen(r)) == 6
			c.memo.starts = append(c.memo.starts, lastStart+uint(found+utf8.RuneLen(r)))
		}
	}
}

func (e *Engine) RmQuery() {
	rmLast(e.query)

	for _, c := range e.cands {
		if c.memo.matched {
			c.memo.starts = rmLast(c.memo.starts)
		}
		c.memo.matched = len(c.memo.starts) == len(e.query)
	}
}

func (e *Engine) DelQuery() {
	if len(e.query) == 0 {
		return
	}
	e.query = []rune{}

	for _, c := range e.cands {
		c.memo.matched = true
		c.memo.starts = []uint{}
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
