package utils

func Max(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func ExtractUniqueValues(array [][]map[string]interface{}) []string{

	uniqueValues := make(map[string]bool)
	var values []string

	for _, maps := range array {
		for _, m := range maps {
			value := m["value"].(string)
			if !uniqueValues[value] {
				uniqueValues[value] = true
				values = append(values, value)
			}
		}
	}
	return values
}
