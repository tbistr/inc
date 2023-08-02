package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tbistr/inc"
)

// NewList creates a new list model for showing candidates.
func NewList(e *inc.Engine) list.Model {
	cands := e.Candidates()
	items := make([]list.Item, len(cands))
	for i, c := range cands {
		items[i] = item{c}
	}
	li := list.New(items, itemDelegate{}, 0, 0)
	li.SetShowTitle(false)
	li.SetShowFilter(false)
	li.SetShowStatusBar(false)
	li.SetShowHelp(false)
	li.KeyMap = list.KeyMap{}

	return li
}

type item struct{ inc.Candidate }

var _ list.Item = item{}

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

var _ list.ItemDelegate = itemDelegate{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	fmt.Fprint(w, printItem(i, index == m.Index(), m.Width()))
}
