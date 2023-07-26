package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tbistr/inc"
)

// RunSlector runs the default selector UI.
func RunSelector(e *inc.Engine) {
	m := NewModel(e)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		panic(err)
	}
}

type Model struct {
	engine   *inc.Engine
	input    textinput.Model
	list     list.Model
	choice   string
	quitting bool
}

func NewModel(e *inc.Engine) Model {
	ti := textinput.New()
	ti.Placeholder = "Enter a query"
	ti.Focus()
	ti.Width = 20

	cands := e.Candidates()
	items := make([]list.Item, len(cands))
	for i, c := range cands {
		items[i] = item{c}
	}
	li := list.New(items, itemDelegate{}, 0, 0)
	li.SetShowTitle(false)
	li.SetShowFilter(false)
	li.SetShowStatusBar(false)
	li.KeyMap = list.KeyMap{}

	return Model{
		engine: e,
		input:  ti,
		list:   li,
	}
}

var _ tea.Model = Model{}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.input.Width = msg.Width
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 5)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.String()
			}
			return m, tea.Quit
		}
	}

	var cmdI, cmdL tea.Cmd
	m.input, cmdI = m.input.Update(msg)
	m.engine.DelQuery()
	for _, r := range m.input.Value() {
		m.engine.AddQuery(r)
	}

	cands := m.engine.Candidates()
	items := []list.Item{}
	for _, c := range cands {
		if c.Matched {
			items = append(items, item{c})
		}
	}
	m.list.SetItems(items)
	m.list, cmdL = m.list.Update(msg)
	return m, tea.Batch(cmdI, cmdL)
}

func (m Model) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s? Sounds good to me.", m.choice)
	}
	if m.quitting {
		return "Not hungry? That’s cool."
	}

	return m.input.View() +
		"\n\n" +
		m.list.View()
}
