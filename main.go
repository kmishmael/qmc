package main

import (
	"fmt"
	"math"
	"qmc/tables"
	"qmc/utils"
	"strconv"
)

func countZeros(binary string) int {
	count := 0
	for _, bit := range binary {
		if bit == '0' {
			count++
		}
	}
	return count
}

func hasCommonBinaryValue(num1, num2 string) (bool, string) {
	diffCount := 0
	combinedValue := ""
	for i := 0; i < len(num1); i++ {
		if num1[i] != num2[i] {
			diffCount++
			if diffCount > 1 {
				return false, ""
			}
			combinedValue += "x"
		} else {
			combinedValue += string(num1[i])
		}
	}
	return diffCount == 1, combinedValue
}

func main() {
	maxterms := []int{0, 1, 3, 4, 5, 6}

	m := utils.Max(maxterms)
	radix := int(math.Log2(float64(m))) + 1

	binTerms := make([]string, len(maxterms))
	for i, term := range maxterms {
		binTerms[i] = strconv.FormatInt(int64(term), 2)
		for len(binTerms[i]) < radix {
			binTerms[i] = "0" + binTerms[i]
		}
	}

	groups := make([][]map[string]interface{}, 0)

	for i := radix; i >= 0; i-- {
		group := make([]map[string]interface{}, 0)
		for _, term := range binTerms {
			c := countZeros(term)
			if c == i {
				group = append(group, map[string]interface{}{"key": []int{toInt(term)}, "value": term, "matched": false})
			}
		}
		if len(group) > 0 {
			groups = append(groups, group)
		}
	}

	primeImplicants := make([]string, 0)
	g := make([][]map[string]interface{}, len(groups))
	copy(g, groups)

	reachedMinReduction := false
	r := 1
	for !reachedMinReduction {
		newGroups := make([][]map[string]interface{}, 0)
		didMatch := false
		for i := 0; i < len(g)-1; i++ {
			group1 := g[i]
			group2 := g[i+1]
			grp := make([]map[string]interface{}, 0)

			for _, term1 := range group1 {
				for _, term2 := range group2 {
					isCommon, newBinary := hasCommonBinaryValue(term1["value"].(string), term2["value"].(string))
					if isCommon {
						term1["matched"] = true
						term2["matched"] = true
						didMatch = true
						newGroup := map[string]interface{}{"key": append(term1["key"].([]int), term2["key"].([]int)...), "value": newBinary, "matched": false}
						grp = append(grp, newGroup)
					}
				}
			}
			newGroups = append(newGroups, grp)
		}
		if !didMatch {
			break
		}
		for _, grp := range g {
			for _, t := range grp {
				if !t["matched"].(bool) {
					primeImplicants = append(primeImplicants, t["value"].(string))
				}
			}
		}
		fmt.Println(len(newGroups), newGroups)
		g = newGroups
		fmt.Println("")
		fmt.Printf("Reduction %d\n", r)
		tables.BuildTable(g)
		r++

	}

	fmt.Println("")
	fmt.Println("Prime implicants: ")
	for _, implicant := range primeImplicants {
		fmt.Printf("=> %s\n", implicant)
	}

}

func toInt(binary string) int {
	val, _ := strconv.ParseInt(binary, 2, 64)
	return int(val)
}
