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
		Width(s.widths.Command).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) ContentStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		Width(s.widths.Content).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) DescriptionStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		Width(s.widths.Description).
		PaddingLeft(1).
		PaddingRight(1)
}

// wrapText wraps text to fit within the given width (in display cells)
func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	var currentLine string

	for _, word := range words {
		// Calculate the display width of the current line with the new word
		testLine := currentLine
		if testLine != "" {
			testLine += " " + word
		} else {
			testLine = word
		}

		displayWidth := runewidth.StringWidth(testLine)
		if displayWidth <= width {
			currentLine = testLine
		} else {
			// Current line is full, save it and start a new one
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			// If a single word is longer than width, break it
			wordWidth := runewidth.StringWidth(word)
			if wordWidth > width {
				// Break the word if it's too long
				runes := []rune(word)
				var currentWord string
				for _, r := range runes {
					testWord := currentWord + string(r)
					if runewidth.StringWidth(testWord) > width {
						if currentWord != "" {
							lines = append(lines, currentWord)
						}
						currentWord = string(r)
					} else {
						currentWord = testWord
					}
				}
				currentLine = currentWord
			} else {
				currentLine = word
			}
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	if len(lines) == 0 {
		return []string{""}
	}

	return lines
}

// commandStyleWithoutWidth returns a style without Width constraint
func (s *Style) commandStyleWithoutWidth() lipgloss.Style {
	return s.lipgloss.
		Foreground(Magenta).
		Background(ContentBg).
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) contentStyleWithoutWidth() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) descriptionStyleWithoutWidth() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		PaddingLeft(1).
		PaddingRight(1)
}

// RenderCell wraps text and renders it within the column width
func (s *Style) RenderCommandCell(text string) []string {
	// Calculate available width: column width minus left and right padding
	availableWidth := s.widths.Command - 2 // 1 left + 1 right padding
	if availableWidth <= 0 {
		availableWidth = 1
	}
	lines := wrapText(text, availableWidth)
	
	style := s.commandStyleWithoutWidth()
	// Use Width() and Align() to ensure consistent column width
	widthStyle := lipgloss.NewStyle().Width(s.widths.Command).Align(lipgloss.Left)
	var rendered []string
	for _, line := range lines {
		renderedLine := style.Render(line)
		// Apply width constraint to ensure consistent column width
		// AlignLeft prevents re-wrapping since the line is already within the width
		rendered = append(rendered, widthStyle.Render(renderedLine))
	}
	return rendered
}

func (s *Style) RenderContentCell(text string) []string {
	availableWidth := s.widths.Content - 2 // 1 left + 1 right padding
	if availableWidth <= 0 {
		availableWidth = 1
	}
	lines := wrapText(text, availableWidth)
	
	style := s.contentStyleWithoutWidth()
	widthStyle := lipgloss.NewStyle().Width(s.widths.Content).Align(lipgloss.Left)
	var rendered []string
	for _, line := range lines {
		renderedLine := style.Render(line)
		rendered = append(rendered, widthStyle.Render(renderedLine))
	}
	return rendered
}

func (s *Style) RenderDescriptionCell(text string) []string {
	availableWidth := s.widths.Description - 2 // 1 left + 1 right padding
	if availableWidth <= 0 {
		availableWidth = 1
	}
	lines := wrapText(text, availableWidth)
	
	style := s.descriptionStyleWithoutWidth()
	widthStyle := lipgloss.NewStyle().Width(s.widths.Description).Align(lipgloss.Left)
	var rendered []string
	for _, line := range lines {
		renderedLine := style.Render(line)
		rendered = append(rendered, widthStyle.Render(renderedLine))
	}
	return rendered
}
