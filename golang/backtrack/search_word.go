package main

import "fmt"

func wordExist(board [][]string, word string) bool {

	row := len(board)
	col := len(board[0])

	visited := make([][]bool, row)
	for i, _ := range visited {
		visited[i] = make([]bool, col)
	}

	var searchWord func(r, c, index int) bool

	searchWord = func(r, c, index int) bool {

		if index == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= row || c >= col || visited[r][c] || board[r][c] != string(word[index]) {
			return false
		}

		visited[r][c] = true

		o1 := searchWord(r, c+1, index+1)
		o2 := searchWord(r, c-1, index+1)
		o3 := searchWord(r+1, c, index+1)
		o4 := searchWord(r-1, c, index+1)

		visited[r][c] = false

		return o1 || o2 || o3 || o4

	}

	var found bool
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			found = searchWord(i, j, 0)
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	return found
}

func main() {

	fmt.Printf("test")

	board := [][]string{
		{"B", "L", "C", "H"},
		{"D", "E", "L", "T"},
		{"D", "A", "K", "A"},
	}

	word := "BLEAK"

	found := wordExist(board, word)

	fmt.Printf("Found %s? %b ", word, found)

}
