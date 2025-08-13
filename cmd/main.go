package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	text string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5"))
	return style.Render(m.text) + "\n\nPress q to Quit."
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func main() {
	initialModel := model{text: "Hello, World!"}
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
