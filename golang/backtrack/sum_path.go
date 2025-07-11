package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

var result [][]int

func backtrack(n *Node, path []int, target int) {

	if n == nil || target < 0 {
		return
	}

	target -= n.Value
	path = append(path, n.Value)

	if target == 0 && n.Left == nil && n.Right == nil {
		result = append(result, path)
	}

	if n.Left != nil {
		backtrack(n.Left, path, target)
	}

	if n.Right != nil {
		backtrack(n.Right, path, target)
	}

}

func main() {

	root := Node{Value: 4}

	n1 := Node{Value: 7}
	n2 := Node{Value: 2}

	root.Left = &n1
	root.Right = &n2

	n3 := Node{Value: 1}
	n4 := Node{Value: 3}
	n5 := Node{Value: 6}
	n6 := Node{Value: 1}

	n1.Left = &n3
	n1.Right = &n4

	n2.Left = &n5
	n2.Right = &n6

	backtrack(&root, []int{}, 7)

	fmt.Println("result is", result)

}
