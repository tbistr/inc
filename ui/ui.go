package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tbistr/inc"
)

// RunSelector runs the default selector UI.
func RunSelector(e *inc.Engine) (bool, inc.Candidate, error) {
	m, err := tea.NewProgram(NewModel(e), tea.WithAltScreen()).Run()
	return m.(Model).canceled, m.(Model).selected, err
}

type keyMap struct {
	Up,
	Down,
	Enter,
	Quit key.Binding
}

var keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up")),
	Down:  key.NewBinding(key.WithKeys("down")),
	Enter: key.NewBinding(key.WithKeys("enter")),
	Quit:  key.NewBinding(key.WithKeys("ctrl+c", "esc")),
}

type Model struct {
	engine   *inc.Engine
	input    textinput.Model
	list     list.Model
	keys     keyMap
	selected inc.Candidate
	canceled bool
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

	return Model{
		engine: e,
		input:  ti,
		list:   NewList(e),
		keys:   keys,
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
		m.list.SetHeight(msg.Height - 1)
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.list.CursorUp()
		case key.Matches(msg, m.keys.Down):
			m.list.CursorDown()

		case key.Matches(msg, m.keys.Enter):
			m.selected = m.engine.Matched()[m.list.Index()]
			return m, tea.Quit
		case key.Matches(msg, m.keys.Quit):
			m.canceled = true
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
	return m.input.View() +
		"\n" +
		m.list.View()
}
