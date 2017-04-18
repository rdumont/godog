package assist

import (
	"github.com/DATA-DOG/godog/gherkin"
)

func ParseInstance(table *gherkin.DataTable) map[string]string {
	instance := map[string]string{}
	for _, row := range table.Rows {
		if len(row.Cells) == 0 {
			continue
		}

		if len(row.Cells) == 1 {
			instance[row.Cells[0].Value] = ""
		} else {
			instance[row.Cells[0].Value] = row.Cells[1].Value
		}
	}
	return instance
}

func ParseList(table *gherkin.DataTable) []map[string]string {
	headRow := table.Rows[0]
	rows := table.Rows[1:]
	header := make([]string, len(headRow.Cells))
	for i, cell := range headRow.Cells {
		header[i] = cell.Value
	}

	list := make([]map[string]string, len(rows))
	for i, row := range rows {
		instance := map[string]string{}
		for j, cell := range row.Cells {
			instance[header[j]] = cell.Value
		}
		list[i] = instance
	}

	return list
}
