package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tbistr/inc"
)

var (
	keyRuneStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
)

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

	str := printItem(i, m.Width())

	fmt.Fprint(w, str)
}
