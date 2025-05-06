package main

import (
	"fmt"
	"math"
)

// Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of
// this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any
// node never differ by more than one.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func isBalancedTree(root *Node, result bool) (int, bool) {

	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := isBalancedTree(root.Left, result)
	rightHeight, rightBalanced := isBalancedTree(root.Right, result)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	height := math.Abs((float64)(leftHeight - rightHeight))
	if height > 1 {
		return 0, false
	}
	rootHeight := max(leftHeight, rightHeight) + 1

	return rootHeight, true

}

func main() {

	T := Tree{}

	T.root = &Node{Val: 1}
	T.root.Left = &Node{Val: 2}
	T.root.Left.Left = &Node{Val: 4}
	T.root.Left.Right = &Node{Val: 5}
	T.root.Right = &Node{Val: 3}

	_, result := isBalancedTree(T.root, true)
	fmt.Println("Is balanced tree -", result)

}
