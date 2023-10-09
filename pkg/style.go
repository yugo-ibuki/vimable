package pkg

import "github.com/charmbracelet/lipgloss"

const (
	Green     = lipgloss.Color("#00FF00")
	Cyan      = lipgloss.Color("#00FFFF")
	Magenta   = lipgloss.Color("#FF00FF")
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
		Foreground(LightBlue).
		Background(lipgloss.Color("#333333"))
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green).
		Background(lipgloss.Color("#000000"))
}

func (s *Style) TableTitleStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Cyan)
}

func (s *Style) TableCellStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Magenta)
}
