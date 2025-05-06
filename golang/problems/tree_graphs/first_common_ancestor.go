package main

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func findFirstCommonAncestorApproach1(p *Node, q *Node) *Node {

	visited := make(map[*Node]bool)

	for p != nil {
		visited[p] = true
		p = p.Parent
	}

	for q != nil {
		if visited[q] {
			return q
		}
		q = q.Parent
	}

	return nil

}

func main() {

	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Left.Parent = root
	root.Right = &Node{Value: 3}
	root.Right.Parent = root

	root.Left.Left = &Node{Value: 4}
	root.Left.Left.Parent = root.Left

	root.Right.Right = &Node{Value: 5}
	root.Right.Right.Parent = root.Right

	result := findFirstCommonAncestorApproach1(root.Left.Left, root.Right.Right)

	fmt.Println("Result is", result.Value)

}
