package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/vimable/pkg"
)

// テーブル名の短縮コードマッピング
var tableShortCodes = map[string]string{
	"n": "ノーマルモード",
	"v": "ビジュアルブロックモード",
	"c": "コマンドモード",
	"f": "ファイル一覧",
	"l": "LSP補完",
	"s": "vim-surround",
	"b": "ブックマーク(マーク)",
	"m": "ブックマーク(マーク)", // 別名として「m」も追加
}

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
  vimable -t "n"             - 短縮コードでテーブルを指定（n=ノーマルモード）
  vimable -t "n,s"           - 複数のテーブルを短縮コードで指定
  vimable -l                 - 利用可能なテーブル一覧と短縮コードを表示
  vimable -h                 - ヘルプを表示`,

	Example: `  # すべてのコマンドを表示
  vimable

  # ノーマルモードとvim-surroundのコマンドのみを表示
  vimable -t "ノーマルモード,vim-surround"
  
  # 短縮コードを使用して同じ表示をする
  vimable -t "n,s"

  # 利用可能なテーブル一覧と短縮コードを表示
  vimable -l`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	// フラグの追加
	rootCmd.Flags().StringVarP(&tableFlag, "table", "t", "", "表示するテーブルを指定します（カンマ区切りで複数指定可能、短縮コード使用可）")
	rootCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "利用可能なテーブル一覧と短縮コードを表示します")
}

// 短縮コードをテーブル名に変換する
func expandTableCode(code string) string {
	if fullName, ok := tableShortCodes[code]; ok {
		return fullName
	}
	return code // 短縮コードが見つからない場合は元のコードを返す
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
		
		// 短縮コードの逆引きマップを作成（テーブル名から短縮コードを取得するため）
		reverseMap := make(map[string][]string)
		for code, name := range tableShortCodes {
			reverseMap[name] = append(reverseMap[name], code)
		}
		
		for key := range data {
			if codes, ok := reverseMap[key]; ok && len(codes) > 0 {
				// 短縮コードがある場合は表示
				fmt.Printf("- %s (短縮コード: %s)\n", key, strings.Join(codes, ", "))
			} else {
				fmt.Printf("- %s\n", key)
			}
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
			// 短縮コードをテーブル名に変換
			expandedTable := expandTableCode(table)
			
			if datums, ok := data[expandedTable]; ok {
				filteredData[expandedTable] = datums
			} else {
				// 短縮コードが変換されている場合は両方の名前を表示
				if expandedTable != table {
					fmt.Printf("警告: テーブル '%s'(%s) は存在しません。\n", table, expandedTable)
				} else {
					fmt.Printf("警告: テーブル '%s' は存在しません。\n", table)
				}
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
			style.CommandStyle().Render(pkg.TruncateText(header[0], columnWidths.Command)),
			style.ContentStyle().Render(pkg.TruncateText(header[1], columnWidths.Content)),
			style.DescriptionStyle().Render(pkg.TruncateText(header[2], columnWidths.Description)),
		}
		fmt.Print(lipgloss.JoinHorizontal(lipgloss.Left, headerCells...))
		fmt.Print("\n")

		// display data
		for _, datum := range datums {
			row := []string{
				style.CommandStyle().Render(pkg.TruncateText(datum.Command, columnWidths.Command)),
				style.ContentStyle().Render(pkg.TruncateText(datum.Content, columnWidths.Content)),
				style.DescriptionStyle().Render(pkg.TruncateText(datum.Description, columnWidths.Description)),
			}
			fmt.Print(lipgloss.JoinHorizontal(lipgloss.Left, row...))
			fmt.Print("\n")
		}
	}

	os.Exit(0)
}
