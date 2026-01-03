package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(Gof{}, tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %+v\n", err)
	}
}
