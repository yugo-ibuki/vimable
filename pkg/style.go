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
		MaxWidth(s.widths.Command).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) ContentStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		Width(s.widths.Content).
		MaxWidth(s.widths.Content).
		PaddingLeft(1).
		PaddingRight(1)
}

func (s *Style) DescriptionStyle() lipgloss.Style {
	return s.lipgloss.
		Foreground(TextColor).
		Background(ContentBg).
		Width(s.widths.Description).
		MaxWidth(s.widths.Description).
		PaddingLeft(1).
		PaddingRight(1)
}

// TruncateText truncates text to fit within the specified width
// If text is longer than width, it will be truncated with "..." at the end
// This function considers the display width (full-width characters count as 2)
func TruncateText(text string, width int) string {
	// 改行、タブ、連続する空白を単一のスペースに置き換える
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.Join(strings.Fields(text), " ")
	
	// パディングを考慮して実際の利用可能幅を計算（左右1文字ずつのパディング）
	availableWidth := width - 2
	if availableWidth <= 0 {
		return text
	}
	
	// 表示幅を計算（全角文字は2、半角文字は1としてカウント）
	textWidth := runewidth.StringWidth(text)
	if textWidth <= availableWidth {
		return text
	}
	
	// "..."の表示幅（3文字）を確保
	if availableWidth <= 3 {
		return "..."
	}
	
	// 利用可能な幅から"..."の分を引く
	targetWidth := availableWidth - 3
	
	// 表示幅に基づいてテキストを切り詰める
	truncated := runewidth.Truncate(text, targetWidth, "")
	return truncated + "..."
}
