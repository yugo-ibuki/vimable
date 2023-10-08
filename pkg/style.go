package pkg

import "github.com/charmbracelet/lipgloss"

const (
	Green = lipgloss.Color("#00FF00")
	Cyan  = lipgloss.Color("#00FFFF")
	Pink  = lipgloss.Color("#FF69B4")
)

type Style struct {
	lipgloss lipgloss.Style
}

func NewStyle(width int) Style {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(width)
	return Style{
		lipgloss: style,
	}
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green) // green
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Cyan) // cyan
}

func (s *Style) TableCellStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Pink) // hot pink
}
