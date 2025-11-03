package pkg

import (
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

type Table struct{}

func NewTable() Table {
	return Table{}
}

type ColumnWidths struct {
	Command     int
	Content     int
	Description int
}

// Width - returns the optimal widths for each column based on terminal width
func (t *Table) Width(header []string, data Data) ColumnWidths {
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || termWidth <= 0 {
		termWidth = 80 // fallback width
	}

	paddingPerColumn := 2 // style padding: left+right
	availableWidth := termWidth - (paddingPerColumn * 3)
	if availableWidth < 3 {
		availableWidth = 3
	}

	maxCommand := measureDisplayWidth(header[0])
	maxContent := measureDisplayWidth(header[1])
	maxDesc := measureDisplayWidth(header[2])

	for _, items := range data {
		for _, item := range items {
			maxCommand = max(maxCommand, measureDisplayWidth(item.Command))
			maxContent = max(maxContent, measureDisplayWidth(item.Content))
			maxDesc = max(maxDesc, measureDisplayWidth(item.Description))
		}
	}

	minCommand := clampMinWidth(maxCommand, measureDisplayWidth(header[0]), 4)
	minContent := clampMinWidth(maxContent, measureDisplayWidth(header[1]), 8)
	minDesc := clampMinWidth(maxDesc, measureDisplayWidth(header[2]), 12)

	widths := []int{
		max(maxCommand, minCommand),
		max(maxContent, minContent),
		max(maxDesc, minDesc),
	}

	total := widths[0] + widths[1] + widths[2]
	limit := availableWidth
	if limit <= 0 {
		limit = total
	}

	minLimits := []int{minCommand, minContent, minDesc}
	reductionOrder := []int{2, 1, 0} // prioritize shrinking description/content first

	if total > limit {
		for total > limit {
			adjusted := false
			for _, idx := range reductionOrder {
				minAllowed := minLimits[idx]
				if total > limit && widths[idx] > minAllowed {
					widths[idx]--
					total--
					adjusted = true
				}
				if total <= limit {
					break
				}
			}
			if adjusted {
				continue
			}

			for _, idx := range reductionOrder {
				if total <= limit {
					break
				}
				if widths[idx] > 1 {
					widths[idx]--
					total--
					adjusted = true
				}
			}
			if !adjusted {
				break
			}
		}
	}

	return ColumnWidths{
		Command:     max(widths[0], 1),
		Content:     max(widths[1], 1),
		Description: max(widths[2], 1),
	}
}

func measureDisplayWidth(value string) int {
	if value == "" {
		return 0
	}

	clean := strings.ReplaceAll(value, "\t", "    ")
	maxWidth := 0
	for _, line := range strings.Split(clean, "\n") {
		lineWidth := runewidth.StringWidth(line)
		if lineWidth > maxWidth {
			maxWidth = lineWidth
		}
	}
	return maxWidth
}

func clampMinWidth(maxValue, headerWidth, suggested int) int {
	minWidth := max(headerWidth, suggested)
	if maxValue > 0 && maxValue < minWidth {
		return maxValue
	}
	return minWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
