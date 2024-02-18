package tables

import (
	"fmt"
)

const (
	redColor   = "\033[31m"
	resetColor = "\033[0m"
)

func BuildTable(data [][]map[string]interface{}) {

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
}
