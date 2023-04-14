package ui

import (
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/tbistr/inc"
)

// printQuery prints the query input line.
func printQuery(s tcell.Screen, e *inc.Engine) {
	prompt := "QUERY> "
	x := setContents(s, 0, 0, prompt, defStyle)
	x = setContents(s, x, 0, e.Query(), defStyle)
	setContents(s, x, 0, "_", defStyle)
}

// printCand prints a candidate line.
func printCand(s tcell.Screen, cand inc.Candidate, y int) {
	t := cand.Text
	w, _ := s.Size()
	lastFound := lastOr(cand.FoundRunes(), inc.FoundRune{})
	start, _ := truncate(t, int(lastFound.Pos), w)

	last := uint(start)
	x := 0
	for _, f := range cand.FoundRunes() {
		if int(f.Pos) < start {
			continue
		}
		x = setContents(s, x, y, t[last:f.Pos], defStyle)
		x = setContents(s, x, y, t[f.Pos:f.Pos+f.Len], emphStyle)
		last = f.Pos + f.Len
	}
	setContents(s, x, y, t[last:], defStyle)
}

// setContents is a helper function to set string contents to the screen.
// It returns the next x position.
func setContents(screen tcell.Screen, x int, y int, str string, style tcell.Style) int {
	w, _ := screen.Size()
	for _, r := range str {
		if w <= x {
			break
		}
		screen.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
	return x
}

// truncate a string by screen width.
//
// It returns the index of the first rune to be printed and the truncated string.
// Set last found rune as the 1/3 of the screen width.
func truncate(s string, lastFound int, width int) (index int, truncated string) {
	index = lastFound
	left := width / 3
	// If the last found rune is near the end of the string, print tail of the string by screen width.
	if runewidth.StringWidth(s[lastFound:]) < width-left {
		t := runewidth.TruncateLeft(s, runewidth.StringWidth(s)-width, "")
		return len(s) - len(t), t
	}
	// Otherwise, set last found rune as the 1/3 of the screen width.
	rfor(s[:lastFound], func(r rune) bool {
		left -= runewidth.RuneWidth(r)
		l := utf8.RuneLen(r)
		if l == -1 {
			l = 1
		}
		index -= l
		return left < 0
	})
	return index, runewidth.Truncate(s[index:], width, "")
}

func rfor(s string, f func(rune) (brk bool)) {
	b := []byte(s)
	last := len(b)
	for 0 < last {
		r, size := utf8.DecodeLastRune(b[:last])
		if f(r) {
			return
		}
		last -= size
	}
}

func lastOr[T any](ts []T, defaultV T) T {
	if len(ts) == 0 {
		return defaultV
	}
	return ts[len(ts)-1]
}
