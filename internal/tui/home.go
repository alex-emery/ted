package tui

import (
	"context"
	"fmt"

	"github.com/aemery-cb/ted/internal/style"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Home struct {
	Posts  []*reddit.Post
	list   list.Model
	client *reddit.SubredditService
	styles *style.Styles
	w      int
	h      int
}

type PostItem struct {
	title string
	desc  string
	id    string
}

func (i PostItem) Title() string       { return i.title }
func (i PostItem) Description() string { return i.desc }
func (i PostItem) FilterValue() string { return i.title }

func NewHome(client *reddit.SubredditService, styles *style.Styles) *Home {

	list := list.NewModel(nil, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Home"
	return &Home{client: client, styles: styles, list: list}

}

func (home Home) Init() tea.Cmd {
	return nil
}

func (home Home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "enter":
			item, ok := home.list.SelectedItem().(PostItem)
			if ok {
				return home, SelectPost(item.id)
			}
		}

	case tea.WindowSizeMsg:
		home.w = msg.Width
		home.h = msg.Height
		home.list.SetSize(msg.Width, msg.Height)

	case SelectHomeMsg:
		home.Posts, _, _ = home.client.HotPosts(context.Background(), "all", &reddit.ListOptions{})
		var items []list.Item
		for _, post := range home.Posts {
			items = append(items, home.newPostItem(post))
		}
		home.list = list.NewModel(items, list.NewDefaultDelegate(), 0, 0)
		home.list.Title = "Home"
		home.list.SetSize(home.w, home.h)
	}
	home.list, cmd = home.list.Update(msg)
	cmds = append(cmds, cmd)

	return home, tea.Batch(cmds...)
}

func (home Home) newPostItem(post *reddit.Post) PostItem {
	style := home.styles.PostThumbnail

	title := style.Render(post.Title)
	body := style.Render(fmt.Sprintf("%s - %ss", post.SubredditName, post.Author))
	return PostItem{title: title, desc: body, id: post.ID}
}

func (home Home) View() string {
	if len(home.list.Items()) == 0 {
		return "Loading"
	}
	return home.list.View()
}
