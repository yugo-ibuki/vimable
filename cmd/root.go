package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/pkg"
)

// フラグ変数
var (
	tableFlag string
	listFlag  bool
)

var rootCmd = &cobra.Command{
	Use:   "vimable",
	Short: "You can check the vim command anytime when you forget it.",
	Long: `vimable は Vim コマンドのチートシートを表示するツールです。

使用例:
  vimable                    - すべてのコマンドテーブルを表示
  vimable -t "ノーマルモード"  - 指定したテーブルのみを表示
  vimable -l                 - 利用可能なテーブル一覧を表示
  vimable -h                 - ヘルプを表示`,

	Example: `  # すべてのコマンドを表示
  vimable

  # ノーマルモードとvim-surroundのコマンドのみを表示
  vimable -t "ノーマルモード,vim-surround"

  # 利用可能なテーブル一覧を表示
  vimable -l`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	// フラグの追加
	rootCmd.Flags().StringVarP(&tableFlag, "table", "t", "", "表示するテーブルを指定します（カンマ区切りで複数指定可能）")
	rootCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "利用可能なテーブル一覧を表示します")
}

func Execute() {
	rootCmd.Execute()
}

func run() {
	header := pkg.Header()
	data := pkg.Commands()

	// 利用可能なテーブル一覧を表示する
	if listFlag {
		fmt.Println("利用可能なテーブル一覧:")
		for key := range data {
			fmt.Printf("- %s\n", key)
		}
		os.Exit(0)
	}

	// テーブルをフィルタリングする
	filteredData := make(pkg.Data)
	if tableFlag != "" {
		// カンマで区切られたテーブル名を分割
		tables := strings.Split(tableFlag, ",")
		for _, table := range tables {
			table = strings.TrimSpace(table)
			if datums, ok := data[table]; ok {
				filteredData[table] = datums
			} else {
				fmt.Printf("警告: テーブル '%s' は存在しません。\n", table)
			}
		}
		// フィルタリングした結果が空の場合
		if len(filteredData) == 0 {
			fmt.Println("指定されたテーブルが見つかりません。利用可能なテーブル一覧を表示するには --list または -l オプションを使用してください。")
			os.Exit(1)
		}
		data = filteredData
	}

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
