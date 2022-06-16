package tui

import (
	"context"

	"github.com/aemery-cb/ted/internal/style"
	"github.com/aemery-cb/ted/internal/util"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Post struct {
	post     *reddit.PostAndComments
	client   *reddit.PostService
	styles   *style.Styles
	viewport viewport.Model
	ready    bool
	width    int
	height   int
}

func NewPost(client *reddit.PostService, styles *style.Styles) *Post {

	return &Post{client: client, styles: styles}

}
func (p Post) Init() tea.Cmd {
	return nil
}

func (p Post) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !p.ready {
			p.viewport = viewport.New(msg.Width, msg.Height)
			p.ready = true

		} else {
			p.viewport.Width = msg.Width
			p.viewport.Height = msg.Height
		}

		p.height = msg.Height
		p.width = msg.Width
	case SelectPostMsg:
		post, _, _ := p.client.Get(context.Background(), msg.Id)
		p.post = post
		p.viewport.SetContent(p.Build())

	}

	p.viewport, cmd = p.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

func (p Post) View() string {

	return p.viewport.View()
}

func (p Post) Build() string {
	if p.post == nil {
		return "Select a post"
	}
	rendered := make([]string, 0)
	rendered = append(rendered, p.styles.PostTitle.Render(p.post.Post.Title))
	rendered = append(rendered, util.MarkdownToText(p.post.Post.Body, p.width))
	rendered = append(rendered, p.post.Post.Author)

	for _, comment := range p.post.Comments {
		rendered = append(rendered, p.BuildCommentSection(comment, 0)...)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rendered...)
}

func (p Post) BuildCommentSection(comment *reddit.Comment, indent int) []string {
	style := p.styles.CommentStyle.MarginLeft(indent)
	rendered := []string{
		style.Render(p.styles.CommentAuthorStyle.Render(comment.Author)),
		p.styles.CommentBodyStyle.MarginLeft(indent).Render(util.MarkdownToText(comment.Body, p.width)),
	}
	for _, reply := range comment.Replies.Comments {
		rendered = append(rendered, p.BuildCommentSection(reply, indent+4)...)
	}

	return rendered

}
