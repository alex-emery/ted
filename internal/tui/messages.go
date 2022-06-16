package tui

import tea "github.com/charmbracelet/bubbletea"

func SelectSubreddit(name string) tea.Cmd {
	return func() tea.Msg {
		return SelectSubredditMsg{
			Name: name,
		}
	}
}

func SelectPost(postid string) tea.Cmd {
	return func() tea.Msg {
		return SelectPostMsg{
			Id: postid,
		}
	}
}

func SelectHome() tea.Cmd {
	return func() tea.Msg {
		return SelectHomeMsg{
			Name: "all",
		}
	}
}

func ShowError(err string) tea.Cmd {
	return func() tea.Msg {
		return ErrorMsg{
			Error: err,
		}
	}
}

type ErrorMsg struct {
	Error string
}
type SelectSubredditMsg struct {
	Name string
}

type SelectPostMsg struct {
	Id string
}

type SelectHomeMsg struct {
	Name string
}
