package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "vimable",
	Short: "You can check the vim command anytime when you forget it.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func Execute() {
	rootCmd.Execute()
}

func run() {
	header := pkg.Header()
	data := pkg.Commands()

	// Calculate optimal column widths
	tableInstance := pkg.NewTable()
	columnWidths := tableInstance.Width(header, data)

	// Initialize styles with column widths
	style := pkg.NewStyle(columnWidths)

	// display content
	for key, datums := range data {
		if key != header[0] { // Add extra space only between sections
			fmt.Println()
		}
		fmt.Println(style.ModeStyle().Render(key))

		// display header
		headerCells := []string{
			style.CommandStyle().Render(header[0]),
			style.ContentStyle().Render(header[1]),
			style.DescriptionStyle().Render(header[2]),
		}
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left, headerCells...))

		// display data
		for _, datum := range datums {
			row := []string{
				style.CommandStyle().Render(datum.Command),
				style.ContentStyle().Render(datum.Content),
				style.DescriptionStyle().Render(datum.Description),
			}
			fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left, row...))
		}
	}

	os.Exit(0)
}
