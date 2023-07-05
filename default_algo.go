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

func (d *DefaultAlgo) AppendCands(cands []*Candidate) {
	d.cands = append(d.cands, cands...)
	// TODO: Care about Append after AddQuery.
}

func (d *DefaultAlgo) GetQuery() []rune {
	return d.query
}

// AddQuery adds a rune to the query.
func (d *DefaultAlgo) AddQuery(r rune) {
	d.query = append(d.query, r)

	for _, c := range d.cands {
		if c.Matched {
			last := lastOr(c.KeyRunes, KeyRune{Pos: 0, Len: 0})
			surplus := c.String()[last.Pos+last.Len:]
			found := strings.IndexRune(surplus, r)
			if found == -1 {
				c.Matched = false
				continue
			}

			// head    surplus
			// "123" + "四五六"
			// if addQuery('四') ->
			// Pos = lPos + lLen + found = 2 + 1 + 0 = 3
			// Len = RuneLen('四') = 3
			c.KeyRunes = append(c.KeyRunes, KeyRune{
				Pos: last.Pos + last.Len + uint(found),
				Len: uint(utf8.RuneLen(r)),
			})
		}
	}
}

// RmQuery removes the last rune from the query.
func (d *DefaultAlgo) RmQuery() {
	d.query = rmLast(d.query)

	for _, c := range d.cands {
		if c.Matched {
			c.KeyRunes = rmLast(c.KeyRunes)
		}
		c.Matched = len(c.KeyRunes) == len(d.query)
	}
}

// DelQuery removes all runes from the query.
// All candidates will be matched.
func (d *DefaultAlgo) DelQuery() {
	if len(d.query) == 0 {
		return
	}
	d.query = []rune{}

	for _, c := range d.cands {
		c.Matched = true
		c.KeyRunes = []KeyRune{}
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
