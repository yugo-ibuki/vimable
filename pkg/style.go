package pkg

import "github.com/charmbracelet/lipgloss"

type Style struct{}

func NewStyle() Style {
	return Style{}
}

func (s *Style) HeaderStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")). // green
		Border(lipgloss.NormalBorder()).
		Bold(true)
	return style
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")). // cyan
		Bold(true)
	return style
}

func (s *Style) TableCellStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")). // hot pink
		Bold(true)
	return style
}
