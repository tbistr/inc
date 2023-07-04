package algorithm

import (
	"strings"
	"unicode/utf8"

	"github.com/tbistr/inc"
)

type Default struct {
	query []rune
	cands []inc.InnerCandidate
}

var _ inc.Algorithm = (*Default)(nil)

func (d *Default) AppendCands(ics []inc.InnerCandidate) {
	d.cands = append(d.cands, ics...)
}

func (d *Default) GetQuery() []rune {
	return d.query
}

// AddQuery adds a rune to the query.
func (d *Default) AddQuery(r rune) {
	d.query = append(d.query, r)

	for _, c := range d.cands {
		if c.Matched() {
			last := lastOr(c.GetKeyRunes(), inc.FoundRune{Pos: 0, Len: 0})
			surplus := c.String()[last.Pos+last.Len:]
			found := strings.IndexRune(surplus, r)
			if found == -1 {
				c.SetMatched(false)
				continue
			}

			// head    surplus
			// "123" + "四五六"
			// if addQuery('四') ->
			// Pos = lPos + lLen + found = 2 + 1 + 0 = 3
			// Len = RuneLen('四') = 3
			founds := append(c.GetKeyRunes(), inc.FoundRune{
				Pos: last.Pos + last.Len + uint(found),
				Len: uint(utf8.RuneLen(r)),
			})
			c.SetKeyRunes(founds)
		}
	}
}

// RmQuery removes the last rune from the query.
func (d *Default) RmQuery() {
	d.query = rmLast(d.query)

	for _, c := range d.cands {
		if c.Matched() {
			c.SetKeyRunes(rmLast(c.GetKeyRunes()))
		}
		c.SetMatched(len(c.GetKeyRunes()) == len(d.query))
	}
}

// DelQuery removes all runes from the query.
// All candidates will be matched.
func (d *Default) DelQuery() {
	if len(d.query) == 0 {
		return
	}
	d.query = []rune{}

	for _, c := range d.cands {
		c.SetMatched(true)
		c.SetKeyRunes([]inc.FoundRune{})
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
