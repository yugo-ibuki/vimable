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
	isFirst := true
	for key, datums := range data {
		if !isFirst {
			fmt.Print("\n") // 最小限の改行
		}
		isFirst = false
		
		style.SetTitle(key) // タイトルを設定
		fmt.Print(style.ModeStyle().Render(key))
		fmt.Print("\n")

		// display header
		headerCells := []string{
			style.CommandStyle().Render(header[0]),
			style.ContentStyle().Render(header[1]),
			style.DescriptionStyle().Render(header[2]),
		}
		fmt.Print(lipgloss.JoinHorizontal(lipgloss.Left, headerCells...))
		fmt.Print("\n")

		// display data
		for _, datum := range datums {
			row := []string{
				style.CommandStyle().Render(datum.Command),
				style.ContentStyle().Render(datum.Content),
				style.DescriptionStyle().Render(datum.Description),
			}
			fmt.Print(lipgloss.JoinHorizontal(lipgloss.Left, row...))
			fmt.Print("\n")
		}
	}

	os.Exit(0)
}
