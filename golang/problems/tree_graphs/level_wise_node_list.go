package main

import "fmt"

// List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes
// at each depth (e.g., if you have a tree with depth D, you'll have D linked lists).
type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) GetLevelWiseNodesDFS(current *Node, level int, result *[][]Node) {

	if current == nil {
		return
	}

	if level >= len(*result) {
		*result = append(*result, []Node{})
	}

	(*result)[level] = append((*result)[level], *current)

	bt.GetLevelWiseNodesDFS(current.left, level+1, result)
	bt.GetLevelWiseNodesDFS(current.right, level+1, result)

}

func (bt *BinaryTree) GetLevelWiseNodesBFS() [][]Node {

	currentLevelNodes := []Node{}
	currentLevelNodes = append(currentLevelNodes, *bt.root)

	result := [][]Node{}

	for len(currentLevelNodes) > 0 {

		result = append(result, currentLevelNodes)

		temp := []Node{}
		for _, node := range currentLevelNodes {
			if node.left != nil {
				temp = append(temp, *node.left)
			}
			if node.right != nil {
				temp = append(temp, *node.right)
			}
		}

		currentLevelNodes = temp

	}

	return result

}

func (bt *BinaryTree) Insert(data int) {

	newNode := &Node{
		data: data,
	}

	if bt.root == nil {
		bt.root = newNode
		return
	}

	current := bt.root
	for current != nil {
		if data <= current.data {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		}
	}

}

func main() {

	binaryTree := BinaryTree{}

	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(2)
	binaryTree.Insert(1)
	binaryTree.Insert(6)

	var result [][]Node

	binaryTree.GetLevelWiseNodesDFS(binaryTree.root, 0, &result)

	bfsResults := binaryTree.GetLevelWiseNodesBFS()

	// fmt.Println(result)
	// Print each level
	for level, nodes := range result {
		fmt.Printf("Level %d: ", level)
		for _, node := range nodes {
			fmt.Printf("%d ", node.data)
		}
		fmt.Println()
	}

	for level, nodes := range bfsResults {
		fmt.Printf("Level %d: ", level)
		for _, node := range nodes {
			fmt.Printf("%d ", node.data)
		}
		fmt.Println()
	}

}
