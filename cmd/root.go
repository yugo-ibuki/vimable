package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/pkg"
	"os"
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
	widths := tableInstance.Width(header, data)

	fmt.Println("tableWidth", widths)

	// style
	style := pkg.NewStyle()

	fmt.Println(
		style.TableTitleStyle(widths.Command).Render(header[0]),
		style.TableTitleStyle(widths.Content).Render(header[1]),
		style.TableTitleStyle(widths.Description).Render(header[2]),
	)

	fmt.Println()
	// データを表示
	for _, datums := range data {
		for _, datum := range datums {
			fmt.Println(
				style.TableCellStyle(widths.Command).Render(datum.Command),
				style.TableCellStyle(widths.Content).Render(datum.Content),
				style.TableCellStyle(widths.Description).Render(datum.Description),
			)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
