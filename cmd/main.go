package main

import (
	"log"
	"todo/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := model.NewModel()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
