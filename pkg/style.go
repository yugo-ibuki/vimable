package pkg

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

const (
	Green     = lipgloss.Color("#a3e2a3")
	Magenta   = lipgloss.Color("#d787d7")
	LightBlue = lipgloss.Color("#7aa2f7")
	Orange    = lipgloss.Color("#ff9e64")
	Yellow    = lipgloss.Color("#e0af68")
	Gray      = lipgloss.Color("#7f849c")
	TextColor = lipgloss.Color("#c8ccd4")

	// 背景色
	PluginBg  = lipgloss.Color("#2c1f1a") // オレンジに合う暗めの背景
	ModeBg    = lipgloss.Color("#2a2522") // 黄色に合う暗めの背景
	ContentBg = lipgloss.Color("#1a1b26") // コンテンツ部分の背景
)

type Style struct {
	lipgloss lipgloss.Style
	widths   ColumnWidths
	title    string // 現在のタイトル
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
	baseStyle := s.lipgloss.
		Padding(0, 1).
		MarginTop(0).
		MarginBottom(0)

	// プラグイン名の場合（vim-surround等）
	if s.isPluginTitle(s.title) {
		return baseStyle.
			Foreground(Orange).
			Background(PluginBg).
			Bold(true)
	}

	// モード名の場合（ノーマルモード等）
	return baseStyle.
		Foreground(Yellow).
		Background(ModeBg).
		Italic(true)
}

// isPluginTitle checks if the title is a plugin name
func (s *Style) isPluginTitle(title string) bool {
	pluginPrefixes := []string{"vim-", "nvim-"}
	for _, prefix := range pluginPrefixes {
		if strings.HasPrefix(strings.ToLower(title), prefix) {
			return true
		}
	}
	return strings.Contains(strings.ToLower(title), "plugin") ||
		strings.Contains(strings.ToLower(title), "surround")
}

// SetTitle sets the current title for styling
func (s *Style) SetTitle(title string) {
	s.title = title
}

func (s *Style) HeaderStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Green).
		Background(ContentBg).
		PaddingLeft(1).
		MarginTop(0).
		MarginBottom(0).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(Gray)
}

func (s *Style) CommandStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(Magenta).
		Background(ContentBg).
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) ContentStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) DescriptionStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) RenderCommand(text string) string {
	return s.CommandStyle().Render(wrapText(text, s.widths.Command))
}

func (s *Style) RenderContent(text string) string {
	return s.ContentStyle().Render(wrapText(text, s.widths.Content))
}

func (s *Style) RenderDescription(text string) string {
	return s.DescriptionStyle().Render(wrapText(text, s.widths.Description))
}

func wrapText(text string, width int) string {
	if width <= 0 {
		return text
	}

	text = strings.ReplaceAll(text, "\t", "    ")

	segments := strings.Split(text, "\n")
	var lines []string
	for _, segment := range segments {
		wrappedLines := wrapLine(segment, width)
		for _, line := range wrappedLines {
			lines = append(lines, padLine(line, width))
		}
	}

	return strings.Join(lines, "\n")
}

func wrapLine(line string, width int) []string {
	if width <= 0 {
		return []string{line}
	}

	if line == "" {
		return []string{""}
	}

	var (
		currentWidth int
		builder      strings.Builder
		result       []string
	)

	for _, r := range line {
		runeWidth := runewidth.RuneWidth(r)
		if currentWidth+runeWidth > width && currentWidth > 0 {
			result = append(result, builder.String())
			builder.Reset()
			currentWidth = 0
		}

		builder.WriteRune(r)
		currentWidth += runeWidth

		if currentWidth >= width {
			result = append(result, builder.String())
			builder.Reset()
			currentWidth = 0
		}
	}

	if builder.Len() > 0 || len(result) == 0 {
		result = append(result, builder.String())
	}

	return result
}

func padLine(line string, width int) string {
	lineWidth := runewidth.StringWidth(line)
	if lineWidth >= width {
		return line
	}
	return line + strings.Repeat(" ", width-lineWidth)
}
