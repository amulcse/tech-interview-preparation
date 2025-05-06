package main

import "fmt"

func main() {
	data := [][]int{
		{1, 3, 5},
		{9, 7, 1},
		{4, 2, 3},
	}

	counterMap := map[int]int{}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			current := data[i][j]
			if _, found := counterMap[current]; !found {
				counterMap[current] = 0
			}
			counterMap[current]++
		}
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			current := data[i][j]
			if counterMap[current] > 1 {
				data[i][j] = 0
			}
		}
	}

	fmt.Println("Data", data)
}
