package model

import (
	"fmt"
	style "todo/internal/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Todo struct {
	Done        bool
	Name        string
	Description string
}

func (t Todo) String() string {
	return t.Name
}

type Model struct {
	Height   int
	Width    int
	Cursor   int
	Selected *Todo
	Todos    []Todo
}

func NewModel() Model {
	return Model{Todos: []Todo{
		{Done: false, Name: "Dishes", Description: "Wash all the dishes from last night."},
		{Done: false, Name: "Buy Groceries", Description: "Stock running low. Need some by next week"},
		{Done: false, Name: "Exercise", Description: "30-reps of cardio stuff."},
	}}
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyUp:
			if m.Cursor > 0 {
				m.Cursor--
			}
		case tea.KeyDown:
			if m.Cursor < len(m.Todos)-1 {
				m.Cursor++
			}
		case tea.KeyEnter:
			m.Selected = &m.Todos[m.Cursor]

		case tea.KeyEsc:
			if m.Selected != nil {
				m.Selected = nil
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var first, second []string

	first = append(first, style.Header("Your To-do List:"))

	if len(m.Todos) == 0 {
		first = append(first, "It's empty.")
	} else {
		for i, todo := range m.Todos {
			shortened := todo.String()

			runes := []rune(shortened)

			if len(runes) > 24 {
				shortened = string(runes[:20]) + "..."
			}

			if i == m.Cursor {
				first = append(first, style.TodoHighlight(fmt.Sprintf(" ► %d.   %s", i+1, shortened)))
			} else {
				first = append(first, style.Todo(fmt.Sprintf("   %d.   %s", i+1, shortened)))
			}
		}
	}

	second = append(second, lipgloss.JoinVertical(lipgloss.Left, first...))

	if m.Selected != nil {
		second = append(second, m.Selected.Description)
	}

	return lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Left, 0.8,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			second...,
		),
	)
}
