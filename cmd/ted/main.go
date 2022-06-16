package main

import (
	"fmt"
	"os"

	"github.com/aemery-cb/ted/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.NewWindow(),
		tea.WithAltScreen(), // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
