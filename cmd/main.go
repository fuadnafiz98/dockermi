package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]bool
}

func initialModel() model {
	return model{
		choices:  []string{"docker image 1", "docker image 2", "docker image 3"},
		selected: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.selected[m.cursor] {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))
	itemStyle := lipgloss.NewStyle()
	selectedItemStyle := lipgloss.NewStyle()
	cursorStyle := lipgloss.NewStyle()

	s := titleStyle.Render("Choose a Docker image:")
	s += "\n"

	for i, choice := range m.choices {
		cursor := ""
		if m.cursor == i {
			cursor = cursorStyle.Render(">")
		}

		checked := " "
		if m.selected[i] {
			checked = "x"
		}

		line := fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

		if m.cursor == i {
			s += selectedItemStyle.Render(line)
		} else {
			s += itemStyle.Render(line)
		}
	}

	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#f0f0f0")).MarginTop(1)
	footer := footerStyle.Render("\nPress q to quit.")
	s += "\n" + footer
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
