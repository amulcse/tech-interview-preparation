package main

import "fmt"

func maskDuplicate(grid [][]int) [][]int {

	rows := len(grid)
	cols := len(grid[0])

	masks := make([][]bool, rows)
	for i := range masks {
		masks[i] = make([]bool, cols)
	}

	// iterate over rows
	for i := 0; i < rows; i++ {
		start := 0
		for j := 1; j <= cols; j++ {

			if j == cols || grid[i][j] != grid[i][start] {
				if j-start+1 >= 3 {
					for k := start; k < j; k++ {
						masks[i][k] = true
					}
				}
				start = j
			}

		}
	}

	for j := 0; j < cols; j++ {
		start := 0
		for i := 1; i <= rows; i++ {

			if i == rows || grid[i][j] != grid[start][j] {

				if i-start+1 >= 3 {
					for k := start; k < i; k++ {
						masks[k][j] = true
					}
				}
				start = i

			}

		}
	}

	fmt.Println("mask", masks)
	for i := range rows {
		for j := range cols {
			if masks[i][j] {
				grid[i][j] = 0
			}
		}
	}

	return grid

}

func main() {

	grid := [][]int{
		{1, 1, 1, 1, 1},
		{4, 2, 1, 4, 1},
	}

	fmt.Println("Remove duplication", maskDuplicate(grid))

}
