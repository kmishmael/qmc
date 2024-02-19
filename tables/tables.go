package tables

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// const (
// 	redColor   = "\033[31m"
// 	resetColor = "\033[0m"
// )

func BuildTable(data [][]map[string]interface{}) {

	// Create a new table
	table := tablewriter.NewWriter(os.Stdout)

	// Extract headers from the first map
	headers := []string{"key", "value", "matched"}

	// Set the table header
	table.SetHeader(headers)

	// Append data to the table
	for _, innerSlice := range data {
		for _, row := range innerSlice {
			var rowValues []string
			for _, header := range headers {
				// Convert interface{} to string based on the type
				val := row[header]
				switch v := val.(type) {
				case int:
					rowValues = append(rowValues, strconv.Itoa(v))
				case string:
					rowValues = append(rowValues, v)
				default:
					// Handle other types if needed
					rowValues = append(rowValues, fmt.Sprintf("%v", v))
				}
			}
			table.Append(rowValues)
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

	/*

		// Output table header
		fmt.Println("| Key | Value | Matched |")
		fmt.Println("| --- | ----- | ------- |")

		// Iterate over data and print rows
		for _, group := range data {
			for _, item := range group {
				key := item["key"]
				value := item["value"]
				matched := item["matched"]

				// Check if "Matched" is false, then apply color
				color := ""
				if !matched.(bool) {
					color = redColor
				}
				// Format output row with color
				fmt.Printf("| %v | %v | %s%v%s |\n", key, value, color, matched, resetColor)

			}

		}
	*/
}
