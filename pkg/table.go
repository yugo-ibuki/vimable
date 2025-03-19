package pkg

import (
	"math"
	"os"

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
	if err != nil {
		termWidth = 80 // fallback width
	}

	// Calculate maximum content lengths for each column
	maxCommand := len(header[0])
	maxContent := len(header[1])
	maxDesc := len(header[2])

	for _, items := range data {
		for _, item := range items {
			maxCommand = max(maxCommand, len(item.Command))
			maxContent = max(maxContent, len(item.Content))
			maxDesc = max(maxDesc, len(item.Description))
		}
	}

	// Add padding and spacing between columns
	spacing := 4 // 2 spaces on each side
	availableWidth := termWidth - (spacing * 3) // space for 3 columns

	// Minimum widths for each column (including padding)
	minCommandWidth := min(maxCommand+spacing, int(float64(availableWidth)*0.2))
	minContentWidth := min(maxContent+spacing, int(float64(availableWidth)*0.3))
	minDescWidth := min(maxDesc+spacing, int(float64(availableWidth)*0.5))

	// Adjust widths if they exceed terminal width
	totalWidth := minCommandWidth + minContentWidth + minDescWidth
	if totalWidth > availableWidth {
		// Scale down proportionally
		ratio := float64(availableWidth) / float64(totalWidth)
		minCommandWidth = int(math.Floor(float64(minCommandWidth) * ratio))
		minContentWidth = int(math.Floor(float64(minContentWidth) * ratio))
		minDescWidth = availableWidth - minCommandWidth - minContentWidth
	}

	return ColumnWidths{
		Command:     minCommandWidth,
		Content:     minContentWidth,
		Description: minDescWidth,
	}
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
