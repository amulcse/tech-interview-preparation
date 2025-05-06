package main

import (
	"fmt"
	"sort"
)

type depthCount struct {
	depth int
	count int
}

func main() {

	s := "(()()())()"

	maxDepth := 0
	depthWiseCount := map[int]int{}
	depth := 0

	for _, c := range s {

		if c == '(' {
			depth++

			if depth > maxDepth {
				maxDepth = depth
			}

		} else if c == ')' {
			depthWiseCount[depth]++
			depth--
		}

	}

	results := []depthCount{}
	for d, c := range depthWiseCount {
		results = append(results, depthCount{depth: d, count: c})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})

	fmt.Println("depth wise count", results)

}
