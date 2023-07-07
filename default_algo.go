package inc

import (
	"strings"
	"unicode/utf8"
)

type DefaultAlgo struct {
	query []rune
	cands []*Candidate
}

var _ Algorithm = (*DefaultAlgo)(nil)

// findAndMark finds runes in the query from the candidate and marks them.
// This starts from the current last key rune.
func findAndMark(query []rune, c *Candidate) {
	lastKey := lastOr(c.KeyRunes, KeyRune{Pos: 0, Len: 0})
	tail := c.String()[lastKey.Pos+lastKey.Len:]

	for _, r := range query {
		found := strings.IndexRune(tail, r)
		if found == -1 {
			c.Matched = false
			return
		}

		// head     tail
		// "123" + "四五六"
		// if findAndMark('四') ->
		// Pos = lPos + lLen + found = 2 + 1 + 0 = 3
		// Len = RuneLen('四') = 3
		lastKey = KeyRune{
			Pos: lastKey.Pos + lastKey.Len + uint(found),
			Len: uint(utf8.RuneLen(r)),
		}
		c.KeyRunes = append(c.KeyRunes, lastKey)
		tail = tail[lastKey.Len:]
	}
}

func (a *DefaultAlgo) AppendCands(cands []*Candidate) {
	for _, c := range cands {
		c.Matched = true
		c.KeyRunes = nil
		findAndMark(a.query, c)
	}
	a.cands = append(a.cands, cands...)
}

func (d *DefaultAlgo) GetQuery() []rune {
	return d.query
}

// AddQuery adds a rune to the query.
func (a *DefaultAlgo) AddQuery(r rune) {
	a.query = append(a.query, r)

	for _, c := range a.cands {
		if c.Matched {
			findAndMark(a.query, c)
		}
	}
}

// RmQuery removes the last rune from the query.
func (a *DefaultAlgo) RmQuery() {
	a.query = rmLast(a.query)

	for _, c := range a.cands {
		if c.Matched {
			c.KeyRunes = rmLast(c.KeyRunes)
		}
		c.Matched = len(c.KeyRunes) == len(a.query)
	}
}

// DelQuery removes all runes from the query.
// All candidates will be matched.
func (a *DefaultAlgo) DelQuery() {
	if len(a.query) == 0 {
		return
	}
	a.query = nil

	for _, c := range a.cands {
		c.Matched = true
		c.KeyRunes = nil
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
