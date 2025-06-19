package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Header        = lipgloss.NewStyle().Padding(1).Bold(true)
	Todo          = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	HighlightTodo = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Bold(true).Background(lipgloss.Color("#ffbf00"))

	Note = lipgloss.NewStyle().Foreground(lipgloss.Color("#3f3f3f"))
)
