package pkg

import "github.com/charmbracelet/lipgloss"

const (
	Green = lipgloss.Color("#00FF00")
	Cyan  = lipgloss.Color("#00FFFF")
	Pink  = lipgloss.Color("#FF69B4")
)

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
		Foreground(Green). // green
		Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(s.Width)
	return style
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(Cyan). // cyan
		Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(s.Width)
	return style
}

func (s *Style) TableCellStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(Pink). // hot pink
		Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(s.Width)
	return style
}
