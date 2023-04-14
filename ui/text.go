package ui

import (
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
	last := uint(0)
	x := 0
	for _, f := range cand.FoundRunes() {
		x = setContents(s, x, y, t[last:f.Pos], defStyle)
		x = setContents(s, x, y, t[f.Pos:f.Pos+f.Len], emphStyle)
		last = f.Pos + f.Len
	}
	setContents(s, x, y, t[last:], defStyle)
}

// setContents is a helper function to set string contents to the screen.
// It treats a runewidth properly.
func setContents(screen tcell.Screen, x int, y int, str string, style tcell.Style) int {
	for _, r := range str {
		screen.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
	return x
}
