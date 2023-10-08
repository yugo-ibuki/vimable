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

// Width - returns the maximum length of the table content.
func (t *Table) Width(header []string, table data.Table) int {
	// header length
	maxLength := headerMax(header)

	for tableTitle, tableData := range table {
		//fmt.Println(tableTitleStyle.Render(tableTitle))
		fmt.Println(tableTitle)

		//fmt.Println(tableCellStyle.Render(command.Command))
		maxInData := tableMax(tableData)
		if maxInData > maxLength {
			maxLength = maxInData
		}
	}

	return maxLength
}

// headerMax - returns the maximum length of the header.
func headerMax(header []string) int {
	headerWidths := make([]int, len(header))
	for i, val := range header {
		headerWidths[i] = len(val)
	}
	return lo.Max(headerWidths)
}

// tableMax - returns the maximum length of the table content.
func tableMax(tableData []data.Data) int {
	dataValue := lo.Map(tableData, func(value data.Data, _ int) []int {
		return []int{
			len(value.Command),
			len(value.Content),
			len(value.Description),
		}
	})
	return lo.Max(lo.Flatten(dataValue))
}
