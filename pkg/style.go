package pkg

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Width int
}

func NewStyle(width int) Style {
	return Style{
		Width: width,
	}
}

func (s *Style) HeaderStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")). // green
		Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(s.Width)
	return style
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")). // cyan
		Bold(true).
		Width(s.Width)
	return style
}

func (s *Style) TableCellStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")). // hot pink
		Bold(true).
		Width(s.Width)
	return style
}
