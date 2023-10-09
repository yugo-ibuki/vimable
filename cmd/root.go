package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/pkg"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "vimable",
	Short: "You can check the vim command anytime when you forget it.",
}

func Execute() {
	header := pkg.Header()
	data := pkg.Commands()

	// table width
	tableInstance := pkg.NewTable()
	tableWidth := tableInstance.Width(header, data)

	fmt.Println("tableWidth", tableWidth)

	// style
	style := pkg.NewStyle(tableWidth)

	// display content
	for key, datums := range data {
		fmt.Println()
		fmt.Println(style.ModeStyle().Render(key))

		// display header
		headerCells := []string{}
		for _, val := range header {
			headerCells = append(headerCells, style.HeaderStyle().Render(val))
		}
		joined := strings.Join(headerCells, " ")
		fmt.Println(lipgloss.JoinHorizontal(0.2, joined))

		// display data
		for _, datum := range datums {
			row := []string{
				style.TableCellStyle().Render(datum.Command),
				style.TableCellStyle().Render(datum.Content),
				style.TableCellStyle().Render(datum.Description),
			}
			fmt.Println(strings.Join(row, " "))
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
