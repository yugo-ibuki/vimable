package pkg

import (
	"github.com/samber/lo"
)

type Table struct{}

func NewTable() Table {
	return Table{}
}

// Width - returns the maximum length of the table content.
func (t *Table) Width(header []string, data Data) int {
	// header length
	maxLength := headerMax(header)

	for _, tableData := range data {
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
func tableMax(tableData []Datum) int {
	dataValue := lo.Map(tableData, func(value Datum, _ int) []int {
		return []int{
			len(value.Command),
			len(value.Content),
			len(value.Description),
		}
	})
	return lo.Max(lo.Flatten(dataValue))
}
