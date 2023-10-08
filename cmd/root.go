package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/data"
	"github.com/yugo-ibuki/vimable/pkg"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "vimable",
	Short: "You can check the vim command anytime when you forget it.",
}

func Execute() {
	table := data.Commands()
	header := data.Header()

	// style
	style := pkg.NewStyle()

	// table width
	tableInstance := pkg.NewTable()
	tableWidth := tableInstance.Width(header, table)

	fmt.Println("tableWidth", tableWidth)

	headerCells := []string{}
	for _, val := range header {
		headerCells = append(headerCells, style.HeaderStyle().Render(val))
	}
	fmt.Println(strings.Join(headerCells, " "))
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
