package pkg

import (
	"github.com/samber/lo"
)

type Table struct{}

type ContentMax struct {
	Command     int
	Content     int
	Description int
}

func NewTable() Table {
	return Table{}
}

// Width - returns the maximum length of the table content.
func (t *Table) Width(header []string, data Data) ContentMax {
	// header length
	headerMaxs := headerMax(header)
	var contentMax ContentMax
	for _, tableData := range data {
		contentMaxs := tableMax(tableData)
		if contentMaxs.Command > headerMaxs.Command {
			contentMax.Command = contentMaxs.Command
		} else {
			contentMax.Command = headerMaxs.Command
		}

		if contentMaxs.Content > headerMaxs.Content {
			contentMax.Content = contentMaxs.Content
		} else {
			contentMax.Content = headerMaxs.Content
		}

		if contentMaxs.Description > headerMaxs.Description {
			contentMax.Description = contentMaxs.Description
		} else {
			contentMax.Description = headerMaxs.Description
		}
	}

	return contentMax
}

// headerMax - returns each length of the header.
func headerMax(headers []string) ContentMax {
	headerWidths := make([]int, len(headers))
	for i, header := range headers {
		headerWidths[i] = len(header)
	}
	return ContentMax{
		headerWidths[0],
		headerWidths[1],
		headerWidths[2],
	}
}

// tableMax - returns the maximum length of the table content.
func tableMax(tableData []Datum) ContentMax {
	commandLengths := lo.Map(tableData, func(value Datum, _ int) int {
		return len(value.Command)
	})
	contentLengths := lo.Map(tableData, func(value Datum, _ int) int {
		return len(value.Content)
	})
	descriptionLengths := lo.Map(tableData, func(value Datum, _ int) int {
		return len(value.Description)
	})
	return ContentMax{
		lo.Max(commandLengths),
		lo.Max(contentLengths),
		lo.Max(descriptionLengths),
	}
}
