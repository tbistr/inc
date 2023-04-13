package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/tbistr/inc"
)

func printQuery(s tcell.Screen, e *inc.Engine) {
	prompt := "QUERY> "
	x := setContents(s, 0, 0, prompt, defStyle)
	x = setContents(s, x, 0, e.Query(), defStyle)
	x = setContents(s, x, 0, "_", defStyle)
	clearSurplus(s, x, 0)
}

func printCand(s tcell.Screen, cand inc.Candidate, y int) {
	t := cand.Text
	last := uint(0)
	x := 0
	for _, f := range cand.FoundRunes() {
		x = setContents(s, x, y, t[last:f.Pos], defStyle)
		x = setContents(s, x, y, t[f.Pos:f.Pos+f.Len], emphStyle)
		last = f.Pos + f.Len
	}
	x = setContents(s, x, y, t[last:], defStyle)
	clearSurplus(s, x, y)
}

func setContents(screen tcell.Screen, x int, y int, str string, style tcell.Style) int {
	for _, r := range str {
		screen.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
	return x
}

func clearSurplus(screen tcell.Screen, x int, y int) {
	w, _ := screen.Size()
	for i := x; i < w; i++ {
		screen.SetContent(i, y, ' ', nil, defStyle)
	}
}
