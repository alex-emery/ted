package tui

import tea "github.com/charmbracelet/bubbletea"

type Error struct {
	err string
}

func (e Error) Init() tea.Cmd {
	return nil
}

func (e Error) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ErrorMsg:
		e.err = msg.Error
	}
	return e, nil
}

func (e Error) View() string {
	return e.err
}
