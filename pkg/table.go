package pkg

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/yugo-ibuki/vimable/data"
)

type Table struct {
}

func NewTable() Table {
	return Table{}
}

func (t *Table) Width(header []string, table data.Table) int {
	// header length
	headerWidths := make([]int, len(header))
	for i, val := range header {
		headerWidths[i] = len(val)
	}

	maxLength := lo.Max(headerWidths)

	for tableTitle, tableData := range table {
		// express table title
		//fmt.Println(tableTitleStyle.Render(tableTitle))
		fmt.Println(tableTitle)

		// express table data
		for _, command := range tableData {
			//fmt.Println(tableCellStyle.Render(command.Command))
			fmt.Println(command.Command)
			dataValue := lo.Map(tableData, func(value data.Data, _ int) []int {
				return []int{
					len(value.Command),
					len(value.Content),
					len(value.Description),
				}
			})
			maxInData := lo.Max(lo.Flatten(dataValue))
			if maxInData > maxLength {
				maxLength = maxInData
			}
		}
	}

	return maxLength
}
