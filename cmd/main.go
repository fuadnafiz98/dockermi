package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fuadnafiz98/dockermi/styles"
	// "dockermi/styles"
)

type model struct {
	counter   int
	quit      bool
	cursor    int
	menuItems []string
}

func initialModel() model {
	return model{
		counter: 0,
		menuItems: []string{
			"Increment",
			"Decrement",
			"Reset",
			"Quit",
		},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if m.quit {
		return "Goodbye!\n"
	}
	var output string

	counterDisplay := styles.DrawDoubleBox("COUNTER", fmt.Sprintf("%d", m.counter), 24)
	menuContent := make([]string, len(m.menuItems))

	for i, item := range m.menuItems {
		if i == m.cursor {
			menuContent[i] = styles.Arrow + " " + item
		} else {
			menuContent[i] = " " + item
		}
	}

	menuBox := styles.DrawBox("MENU", menuContent, 24, m.cursor)

	output = counterDisplay + "\n\n" + menuBox + "\n\n"
	output += "Use ↑/↓ to navigate, Enter to select, q to quit"
	return output
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.cursor {
			case 0:
				m.counter++
			case 1:
				m.counter--
			case 2:
				m.counter = 0
			case 3:
				m.quit = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
