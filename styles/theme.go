package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	TopLeft     = "┌"
	TopRight    = "┐"
	BottomLeft  = "└"
	BottomRight = "┘"
	Horizontal  = "─"
	Vertical    = "│"
	Cross       = "┼"
	TeeUp       = "┴"
	TeeDown     = "┬"
	TeeLeft     = "┤"
	TeeRight    = "├"

	DTopLeft     = "╔"
	DTopRight    = "╗"
	DBottomLeft  = "╚"
	DBottomRight = "╝"
	DHorizontal  = "═"
	DVertical    = "║"

	Arrow      = "▶"
	ArrowEmpty = "▷"
	Block      = "█"
	BlockEmpty = "░"
	Check      = "■"
	Uncheck    = "□"
)

var SharpBorder = lipgloss.Border{
	Top:         Horizontal,
	Bottom:      Horizontal,
	Left:        Vertical,
	Right:       Vertical,
	TopLeft:     TopLeft,
	TopRight:    TopRight,
	BottomLeft:  BottomLeft,
	BottomRight: BottomRight,
}

var (
	BoxStyle = lipgloss.NewStyle().
			Border(SharpBorder).
			BorderForeground(lipgloss.Color("255")).
			Padding(0, 1)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color("255")).
			Bold(true)

	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255"))

	DimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Border(SharpBorder).
			BorderForeground(lipgloss.Color("255")).
			Align(lipgloss.Center).
			Width(40)
)

func RepeatStr(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func PadRight(s string, n int) string {
	if len(s) >= n {
		return s[:n]
	}
	return s + RepeatStr(" ", n-len(s))
}

func PadCenter(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	leftPad := (width - len(s)) / 2
	rightPad := width - len(s) - leftPad
	return RepeatStr(" ", leftPad) + s + RepeatStr(" ", rightPad)
}

func DrawBox(title string, content []string, width int, selected int) string {
	var b strings.Builder

	b.WriteString(TopLeft + RepeatStr(Horizontal, width-2) + TopRight + "\n")

	if title != "" {
		b.WriteString(Vertical + PadCenter(title, width-2) + Vertical + "\n")
		b.WriteString(TeeRight + RepeatStr(Horizontal, width-2) + TeeLeft + "\n")
	}

	for i, line := range content {
		if i == selected {
			b.WriteString(Vertical + "\x1b[7m" + PadRight(line, width-2) + "\x1b[0m" + Vertical + "\n")
		} else {
			b.WriteString(Vertical + PadRight(line, width-2) + Vertical + "\n")
		}
	}

	b.WriteString(BottomLeft + RepeatStr(Horizontal, width-2) + BottomRight)

	return b.String()
}

func DrawDoubleBox(title string, content string, width int) string {
	var b strings.Builder

	b.WriteString(DTopLeft + RepeatStr(DHorizontal, width-2) + DTopRight + "\n")
	b.WriteString(DVertical + PadCenter(title, width-2) + DVertical + "\n")
	b.WriteString(DVertical + PadCenter(content, width-2) + DVertical + "\n")
	b.WriteString(DBottomLeft + RepeatStr(DHorizontal, width-2) + DBottomRight)

	return b.String()
}
