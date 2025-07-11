package main

import "fmt"

type Node struct {
	Value  int
	Childs []*Node
}

var visited map[*Node]*Node

func generateNode(val int) *Node {
	return &Node{Value: val, Childs: []*Node{}}
}

func cloneTree(n *Node) *Node {

	if n == nil {
		return nil
	}

	if val, found := visited[n]; found {
		return val
	}

	clone := generateNode(n.Value)
	visited[n] = clone

	for _, c := range n.Childs {
		clone.Childs = append(clone.Childs, cloneTree(c))
	}

	return clone

}

func traverseTreeLevelBy(root *Node) {

	level := 1
	queue := []*Node{}
	visitedNodes := map[*Node]bool{}

	queue = append(queue, root)

	for len(queue) > 0 {

		nextQueue := []*Node{}
		fmt.Printf("Level %d:", level)

		for _, n := range queue {

			visitedNodes[n] = true
			fmt.Println("current node", n.Value)

			for _, nn := range n.Childs {
				if !visitedNodes[nn] {
					nextQueue = append(nextQueue, nn)
				}
			}

		}

		level++
		queue = nextQueue

	}

}

func main() {

	visited = map[*Node]*Node{}
	n3 := generateNode(3)
	n5 := generateNode(5)

	n1 := generateNode(1)
	n7 := generateNode(7)
	n15 := generateNode(15)
	n20 := generateNode(20)

	n4 := generateNode(4)
	n9 := generateNode(9)

	n3.Childs = append(n3.Childs, n1)
	n3.Childs = append(n3.Childs, n7)

	n1.Childs = append(n1.Childs, n15)
	n7.Childs = append(n7.Childs, n20)

	n5.Childs = append(n5.Childs, n4)

	n4.Childs = append(n4.Childs, n20)
	n4.Childs = append(n4.Childs, n9)

	clone := cloneTree(n3)

	traverseTreeLevelBy(clone)

}
