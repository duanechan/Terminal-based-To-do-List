package style

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Header(text string) string {
	return lipgloss.NewStyle().
		Padding(1).
		Bold(true).
		Render(text)
}

func TodoHighlight(todo string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Bold(true).
		Background(lipgloss.Color("#ffbf00")).
		Render(fmt.Sprintf("%-32s", todo))
}

func Todo(todo string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Render(fmt.Sprintf("%-32s", todo))
}
