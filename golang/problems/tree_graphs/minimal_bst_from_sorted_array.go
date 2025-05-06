package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) insert(data int) {

	fmt.Println("insert data", data)

	newNode := Node{
		data: data,
	}

	if bst.root == nil {
		bst.root = &newNode
		return
	}

	current := bst.root
	previous := current
	for current != nil {
		previous = current
		if data > current.data {
			current = current.right
		} else {
			current = current.left
		}
	}

	if data > previous.data {
		previous.right = &newNode
	} else {
		previous.left = &newNode
	}

}

func (bst *BST) findMiddlePoint(data []int) {

	if len(data) == 0 {
		return
	}

	mid := len(data) / 2
	bst.insert(data[mid])
	bst.findMiddlePoint(data[0:mid])
	bst.findMiddlePoint(data[mid+1:])

}

func main() {

	bst := BST{}

	sortedData := []int{1, 2, 3, 4, 5, 6}

	bst.findMiddlePoint(sortedData)
}
