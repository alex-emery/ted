package style

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	// Posts
	PostTitle          lipgloss.Style
	PostThumbnail      lipgloss.Style
	PostThumbnailHover lipgloss.Style
	// Comments
	CommentStyle       lipgloss.Style
	CommentAuthorStyle lipgloss.Style
	CommentBodyStyle   lipgloss.Style
}

func NewDefaultStyles() *Styles {

	postThumbnail := lipgloss.NewStyle()
	return &Styles{
		PostTitle:          lipgloss.NewStyle().MarginBottom(1),
		PostThumbnail:      postThumbnail,
		PostThumbnailHover: postThumbnail.Copy().Border(lipgloss.NormalBorder(), false, false, false, true).MarginLeft(1),
		CommentStyle:       lipgloss.NewStyle(),
		CommentAuthorStyle: lipgloss.NewStyle().Bold(true).MarginBottom(1),
		CommentBodyStyle:   lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).MarginBottom(1),
	}
}
