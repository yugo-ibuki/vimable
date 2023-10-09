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

func NewStyle() Style {
	style := lipgloss.NewStyle().
		// this is just a default border, but the layout is slightly mispositioned
		//Border(lipgloss.NormalBorder()).
		Bold(true).Padding(0, 1)
	return Style{
		lipgloss: style,
	}
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green) // green
}

func (s *Style) TableTitleStyle(width int) lipgloss.Style {
	return s.lipgloss.
		Foreground(Cyan). // cyan
		Width(width)
}

func (s *Style) TableCellStyle(width int) lipgloss.Style {
	return s.lipgloss.
		Foreground(Pink). // hot pink
		Width(width)
}
