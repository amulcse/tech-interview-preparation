package main

import "fmt"

type Node struct {
	value     int
	neighbors []*Node
}

func (N *Node) addChild(child *Node) {
	N.neighbors = append(N.neighbors, child)
}

func (N *Node) traverse() {

	visited := make(map[*Node]bool)
	queue := []*Node{N}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		fmt.Printf("%d ", current.value)

		for _, n := range current.neighbors {
			queue = append(queue, n)
		}
	}

}

func main() {

	root := Node{value: 2}
	child1 := Node{value: 4}
	child2 := Node{value: 6}
	child3 := Node{value: 8}
	root.addChild(&child1)
	root.addChild(&child2)
	root.addChild(&child3)

	child4 := Node{value: 11}
	child5 := Node{value: 10}
	child1.addChild(&child5)
	child2.addChild(&child4)

	root.traverse()

}
