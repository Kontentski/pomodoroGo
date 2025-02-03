package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	models "github.com/kontentski/pomodoroGo/internal/app"
)

func main() {
	p := tea.NewProgram(
		models.InitialModel(),
		tea.WithAltScreen(), // hides the terminal text

	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
