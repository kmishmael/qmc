package tables

import (
	"fmt"
	"os"
	"strconv"
	"qmc/utils"
	"github.com/olekukonko/tablewriter"
)


func BuildTable(data [][]map[string]interface{}) {

	table := tablewriter.NewWriter(os.Stdout)

	headers := []string{"key", "value", "matched"}

	table.SetHeader(headers)

	for _, innerSlice := range data {
		for _, row := range innerSlice {
			var rowValues []string
			for _, header := range headers {
				val := row[header]
				switch v := val.(type) {
				case int:
					rowValues = append(rowValues, strconv.Itoa(v))
				case string:
					rowValues = append(rowValues, v)
				case []int:
					rowValues = append(rowValues, utils.SliceToSpread(v))
				default:
					rowValues = append(rowValues, fmt.Sprintf("%v", v))
				}
			}

			if matched, ok := row["matched"].(bool); ok && !matched {
				colors := make([]tablewriter.Colors, len(rowValues))
				for i := range colors {
					colors[i] = tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor}
				}

				table.Rich(rowValues, colors)
			} else {
				table.Append(rowValues)
			}
		}
	}

	// Set table format
	table.SetBorder(true)
	table.SetColumnSeparator("|")
	table.SetCenterSeparator("|")
	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Render the table
	table.Render()
}
