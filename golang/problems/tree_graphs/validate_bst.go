package main

// Validate BST: Implement a function to check if a binary tree is a binary search tree.

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func validateBinaryTree(n *Node, lastVisitedNode **Node) bool {

	if n == nil {
		return true
	}

	validLeftTree := validateBinaryTree(n.Left, lastVisitedNode)
	if !validLeftTree {
		return false
	}

	if *lastVisitedNode != nil && (*lastVisitedNode).Value > n.Value {
		return false
	}
	*lastVisitedNode = n

	validRightTree := validateBinaryTree(n.Right, lastVisitedNode)
	if !validRightTree {
		return false
	}

	return true

}

func main() {

	bt := BinaryTree{}

	bt.Root = &Node{Value: 10}
	bt.Root.Left = &Node{Value: 5}
	bt.Root.Right = &Node{Value: 16}
	bt.Root.Left.Left = &Node{Value: 4}
	bt.Root.Left.Right = &Node{Value: 9}

	var lastVisitedNode *Node
	fmt.Println("Is it binary search tree?", validateBinaryTree(bt.Root, &lastVisitedNode))

}
