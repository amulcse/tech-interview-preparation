package main

import (
	"fmt"
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// Helper to check bounds
func inBounds(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func findWords(board [][]rune, dictionary map[string]bool) []string {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	found := make(map[string]bool)

	var dfs func(x, y int, prefix string)
	dfs = func(x, y int, prefix string) {
		if !inBounds(x, y, rows, cols) || visited[x][y] {
			return
		}

		prefix += string(board[x][y])
		if !hasPrefix(dictionary, prefix) {
			return
		}

		if dictionary[prefix] {
			found[prefix] = true
		}

		visited[x][y] = true
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			dfs(nx, ny, prefix)
		}
		visited[x][y] = false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, "")
		}
	}

	// Convert found map to slice
	words := []string{}
	for word := range found {
		words = append(words, word)
	}
	return words
}

// hasPrefix checks if any word in the dictionary starts with the given prefix
func hasPrefix(dict map[string]bool, prefix string) bool {
	for word := range dict {
		if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

func main() {
	board := [][]rune{
		{'T', 'E', 'S', 'T'},
		{'A', 'K', 'R', 'N'},
		{'C', 'O', 'D', 'E'},
		{'L', 'I', 'P', 'Q'},
	}

	// Example dictionary
	dict := map[string]bool{
		"TEST": true,
		"TAKE": true,
		"CODE": true,
		"TOP":  true,
		"TEN":  true,
		"RIP":  true,
		"INK":  true,
		"TIP":  true,
	}

	words := findWords(board, dict)
	fmt.Println("Found words:")
	for _, w := range words {
		fmt.Println("-", w)
	}
}
