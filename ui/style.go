package ui

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
	"github.com/tbistr/inc"
)

var (
	cursor    = lipgloss.NewStyle().Foreground(lipgloss.Color("#f490f4")).Render("> ")
	itemStyle = lipgloss.NewStyle().PaddingLeft(lipgloss.Width(cursor))

	normalStyle  = lipgloss.NewStyle()
	keyRuneStyle = normalStyle.Copy().Foreground(lipgloss.Color("#28a078"))

	selectedStyle        = normalStyle.Copy().Background(lipgloss.Color("#f490f4"))
	selectedKeyRuneStyle = selectedStyle.Copy().Bold(true)
)

// printItem prints a item line.
func printItem(i item, selected bool, maxWidth int) string {
	t := i.String()
	result := []string{}
	keyRunes := i.KeyRunes
	lastFound := lastOr(keyRunes, inc.KeyRune{})
	start, _ := truncate(t, int(lastFound.Pos), maxWidth)
	last := uint(start)

	style := normalStyle
	keyStyle := keyRuneStyle
	if selected {
		style = selectedStyle
		keyStyle = selectedKeyRuneStyle
	}

	for _, k := range keyRunes {
		if int(k.Pos) < start {
			continue
		}

		result = append(
			result,
			style.Render(t[last:k.Pos]),
			keyStyle.Render(t[k.Pos:k.Pos+k.Len]),
		)

		last = k.Pos + k.Len
	}
	result = append(result, style.Render(t[last:]))

	if selected {
		return cursor + strings.Join(result, "")
	}
	return itemStyle.Render(strings.Join(result, ""))
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
