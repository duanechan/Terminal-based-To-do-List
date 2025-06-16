package main

import (
	"log"
	"todo/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
