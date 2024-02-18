package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	maxterms := []int{0, 1, 3, 4, 5, 6}

	var m int = max(maxterms)

	radix := int(math.Log2(float64(m))) + 1

	bin_terms := make([]string, len(maxterms))

	for i, term := range maxterms {
		bin_terms[i] = strconv.FormatInt(int64(term), 2)
		for len(bin_terms[i]) < radix {
			bin_terms[i] = "0" + bin_terms[i]
		}
	}
	fmt.Println("Hello")
}

func max(terms []int) int {
	var m int = terms[0]
	for _, term := range terms {
		if term > m {
			m = term
		}
	}
	return m
}
