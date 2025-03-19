package pkg

import "github.com/charmbracelet/lipgloss"

const (
	Green     = lipgloss.Color("#a3e2a3")
	Magenta   = lipgloss.Color("#d787d7")
	LightBlue = lipgloss.Color("#7aa2f7")
	Gray      = lipgloss.Color("#7f849c")
	TextColor = lipgloss.Color("#c8ccd4")
)

type Style struct {
	lipgloss lipgloss.Style
	widths   ColumnWidths
}

func NewStyle(widths ColumnWidths) Style {
	style := lipgloss.NewStyle().
		Bold(true)
	return Style{
		lipgloss: style,
		widths:   widths,
	}
}

func (s *Style) ModeStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(LightBlue).
		Background(lipgloss.Color("#1a1b26")).
		Padding(0, 1).
		MarginBottom(1)
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green).
		Background(lipgloss.Color("#16161e")).
		PaddingLeft(1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(Gray)
}

func (s *Style) CommandStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Magenta).
		Bold(true).
		Width(s.widths.Command).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) ContentStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Width(s.widths.Content).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) DescriptionStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Width(s.widths.Description).
		PaddingLeft(1).
		PaddingRight(1)
}
