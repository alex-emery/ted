package tui

import (
	"log"
	"os"

	"github.com/aemery-cb/ted/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type SessionState int

const (
	subredditState SessionState = iota // home page
	postState                          // viewing an individual post
	errorState
)

type Window struct {
	state  SessionState
	client *reddit.Client
	boxes  []tea.Model
}

func NewWindow() *Window {
	if v := os.Getenv("GO_REDDIT_CLIENT_ID"); v == "" {
		log.Fatal("GO_REDDIT_CLIENT_ID is not set")
	}
	if v := os.Getenv("GO_REDDIT_CLIENT_SECRET"); v == "" {
		log.Fatal("GO_REDDIT_CLIENT_ID is not set")
	}
	if v := os.Getenv("GO_REDDIT_CLIENT_USERNAME"); v == "" {
		log.Fatal("GO_REDDIT_CLIENT_ID is not set")
	}
	if v := os.Getenv("GO_REDDIT_CLIENT_PASSWORD"); v == "" {
		log.Fatal("GO_REDDIT_CLIENT_ID is not set")
	}

	client, err := reddit.NewClient(reddit.Credentials{}, reddit.FromEnv)

	if err != nil {
		log.Panic(err)
	}

	styles := style.NewDefaultStyles()
	boxes := make([]tea.Model, 3)

	boxes[subredditState] = NewHome(client.Subreddit, styles)

	boxes[postState] = NewPost(client.Post, styles)
	boxes[errorState] = &Error{}

	return &Window{client: client, boxes: boxes}
}

func (w Window) Init() tea.Cmd {
	return SelectHome()
}

func (w Window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return w, tea.Quit

		case "backspace":
			if w.state == postState {
				w.state = subredditState
			}
		}

	case ErrorMsg:
		w.state = errorState
	case SelectPostMsg:
		w.state = postState
	}

	_, keyMsg := msg.(tea.KeyMsg)
	_, mouseMsg := msg.(tea.MouseMsg)
	if keyMsg || mouseMsg { // only send keyboard to visible view
		newBx, cmd := w.boxes[w.state].Update(msg)
		cmds = append(cmds, cmd)
		w.boxes[w.state] = newBx
	} else {
		for i, bx := range w.boxes {
			newBx, cmd := bx.Update(msg)
			w.boxes[i] = newBx
			cmds = append(cmds, cmd)
		}
	}

	return w, tea.Batch(cmds...)
}

func (w Window) View() string {

	return w.boxes[w.state].View()
}
