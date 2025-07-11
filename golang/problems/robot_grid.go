package main

import "fmt"

var memo [][]int

func numberOfWays(grid [][]int, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])

	if r >= rows || c >= cols || grid[r][c] == -1 {
		return 0
	}

	if memo[r][c] != -1 {
		return memo[r][c]
	}

	if r == rows-1 && c == cols-1 {
		return 1
	}

	memo[r][c] = numberOfWays(grid, r+1, c) + numberOfWays(grid, r, c+1)
	return memo[r][c]
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{-1, 0, 0},
		{0, 0, 0},
	}

	rows := len(grid)
	cols := len(grid[0])

	// Properly initialize memo
	memo = make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	result := numberOfWays(grid, 0, 0)
	fmt.Println("Number of ways:", result)
}
