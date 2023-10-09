package pkg

import "github.com/charmbracelet/lipgloss"

const (
	Green     = lipgloss.Color("#00FF00")
	Cyan      = lipgloss.Color("#00FFFF")
	Pink      = lipgloss.Color("#FF69B4")
	LightBlue = lipgloss.Color("#ADD8E6")
)

type Style struct {
	lipgloss lipgloss.Style
}

func NewStyle(width int) Style {
	style := lipgloss.NewStyle().
		// this is just a default border, but the layout is slightly mispositioned
		//Border(lipgloss.NormalBorder()).
		Bold(true).
		Width(width)
	return Style{
		lipgloss: style,
	}
}

func (s *Style) ModeStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(LightBlue)
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green)
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Cyan)
}

func (s *Style) TableCellStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Pink)
}
