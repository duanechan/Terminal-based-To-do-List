package model

import (
	"fmt"
	"slices"
	"time"
	"todo/internal/style"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Todo struct {
	Id          int
	CreatedAt   int64
	Done        bool
	Name        string
	Description string
}

func (t Todo) String() string {
	return t.Name
}

type Model struct {
	InsertMode bool
	Height     int
	Width      int
	Cursor     int
	FieldIdx   int
	Selected   *Todo
	Todos      []Todo
	Fields     []textinput.Model
}

func NewModel() Model {
	var fields []textinput.Model
	var ti textinput.Model

	for i := range 2 {
		ti = textinput.New()
		ti.Width = 150
		ti.TextStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffbf00"))
		switch i {
		case 0:
			ti.Prompt = "Name: "
			ti.Placeholder = "Study for Exams"
			ti.CharLimit = 15

		case 1:
			ti.Prompt = "Description (optional): "
			ti.Placeholder = "Do a 30-minute review on data structures & algorithms."
			ti.CharLimit = 64

		}
		fields = append(fields, ti)
	}

	return Model{
		Fields: fields,
		Todos:  []Todo{},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyUp:
			if m.InsertMode && m.FieldIdx > 0 {
				m.Fields[m.FieldIdx].Blur()
				m.FieldIdx--
				m.Fields[m.FieldIdx].Focus()
			}

			if m.Cursor > 0 {
				m.Cursor--
				m.Selected = &m.Todos[m.Cursor]
			}

		case tea.KeyDown:
			if m.InsertMode && m.FieldIdx < len(m.Fields)-1 {
				m.Fields[m.FieldIdx].Blur()
				m.FieldIdx++
				m.Fields[m.FieldIdx].Focus()
			}

			if m.Cursor < len(m.Todos)-1 {
				m.Cursor++
				m.Selected = &m.Todos[m.Cursor]
			}

		case tea.KeyEnter:
			if m.InsertMode {
				if m.FieldIdx < len(m.Fields)-1 {
					m.Fields[m.FieldIdx].Blur()
					m.FieldIdx++
					m.Fields[m.FieldIdx].Focus()
					return m, nil
				} else {
					if m.Fields[0].Value() == "" {
						return m, nil
					}
					todo := Todo{
						Done:        false,
						CreatedAt:   time.Now().Unix(),
						Name:        m.Fields[0].Value(),
						Description: m.Fields[1].Value(),
					}
					m.Todos = append(m.Todos, todo)

					for i := range m.Fields {
						m.Fields[i].SetValue("")
						m.Fields[i].Blur()
					}
					m.FieldIdx = 0
					m.Fields[0].Focus()
					m.InsertMode = false
					m.Cursor = slices.Index(m.Todos, todo)
					m.Selected = &todo

					return m, nil
				}
			}

		case tea.KeyCtrlA:
			m.InsertMode = true
			m.Fields[0].Focus()

		case tea.KeyEsc:
			if m.InsertMode {
				m.InsertMode = false
				m.Fields[0].Blur()
				m.Fields[1].Blur()
				return m, nil
			}
		}
	}

	m.Fields[m.FieldIdx], cmd = m.Fields[m.FieldIdx].Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.InsertMode {
		sections := []string{"Insert Mode"}

		for i := range m.Fields {
			sections = append(sections, m.Fields[i].View())
		}

		return lipgloss.Place(
			m.Width, m.Height,
			lipgloss.Left, 0,
			lipgloss.JoinVertical(
				lipgloss.Left,
				sections...,
			),
		)
	}

	var first, second []string

	first = append(first, style.Header.Render("Your To-do List:"))

	if len(m.Todos) == 0 {
		first = append(first, "It's empty. Press Insert to add an item.")
	} else {
		for i, todo := range m.Todos {
			if i == m.Cursor {
				first = append(first, style.HighlightTodo.Render(fmt.Sprintf(" ► %d.   %-32s", i+1, todo)))
			} else {
				first = append(first, style.Todo.Render(fmt.Sprintf("   %d.   %-32s", i+1, todo)))
			}
		}
	}

	second = append(second, lipgloss.JoinVertical(lipgloss.Left, first...))

	if m.Selected != nil {
		t := time.Unix(m.Selected.CreatedAt, 0)

		second = append(second,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(
				lipgloss.Place(
					m.Width/4, m.Height/2,
					lipgloss.Left, lipgloss.Left,
					lipgloss.JoinVertical(
						lipgloss.Left,
						t.Format("January 2, 2006 3:04 PM")+"\n",
						m.Selected.Name+"\n",
						m.Selected.Description,
					),
				),
			),
		)
	}

	return lipgloss.Place(
		m.Width, m.Height,
		lipgloss.Left, 0,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			second...,
		),
	)
}
