package main

import "fmt"

type TrieNode struct {
	children map[string]*TrieNode
	word     string
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) InsertWord(word string) {

	node := t.root
	for _, c := range word {
		ch := string(c)
		if _, exist := node.children[string(ch)]; !exist {
			node.children[ch] = &TrieNode{children: make(map[string]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.word = word

}

func (t *Trie) DFSOnBoard(board [][]byte, row, col int, node *TrieNode, founds *map[string]bool) {

	rows, cols := len(board), len(board[0])

	if row < 0 || col < 0 || row >= rows || col >= cols || board[row][col] == '#' {
		return
	}

	ch := string(board[row][col])
	oldCh := board[row][col]

	nextNode, exists := node.children[ch]
	if !exists {
		return
	}

	board[row][col] = '#'

	// Check if we found a word
	if nextNode.word != "" {
		(*founds)[nextNode.word] = true
	}

	t.DFSOnBoard(board, row-1, col, nextNode, founds)
	t.DFSOnBoard(board, row, col-1, nextNode, founds)
	t.DFSOnBoard(board, row+1, col, nextNode, founds)
	t.DFSOnBoard(board, row, col+1, nextNode, founds)

	board[row][col] = oldCh

}

func main() {

	words := []string{"oath", "pea", "eat", "rain"}
	t := &Trie{}
	t.root = &TrieNode{children: make(map[string]*TrieNode)}

	founds := map[string]bool{}
	for _, word := range words {
		t.InsertWord(word)
		founds[word] = false
	}

	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 't', 'l', 'v'},
	}

	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			t.DFSOnBoard(board, i, j, t.root, &founds)
		}
	}

	fmt.Println("found match", founds)

}
