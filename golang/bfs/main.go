package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func bfsTraversal(n *Node) {

	queue := []*Node{}
	queue = append(queue, n)

	visible := map[*Node]bool{}

	for len(queue) > 0 {

		c := queue[0]
		queue = queue[1:]

		fmt.Println("Node", c.Value)
		visible[c] = true

		if c.Left != nil && !visible[c.Left] {
			queue = append(queue, c.Left)
		}

		if c.Right != nil && !visible[c.Right] {
			queue = append(queue, c.Right)
		}

	}

}

func levelOrderSum(n *Node) {

	queue := []*Node{}
	queue = append(queue, n)

	levelSum := map[int]int{}
	level := 1

	for len(queue) > 0 {

		currentLevel := []*Node{}
		sum := 0

		for _, c := range queue {

			sum += c.Value

			if c.Left != nil {
				currentLevel = append(currentLevel, c.Left)
			}

			if c.Right != nil {
				currentLevel = append(currentLevel, c.Right)
			}

		}

		levelSum[level] = sum
		queue = currentLevel
		level++

	}

	fmt.Println(levelSum)

}

func rightMostNode(n *Node) {

	queue := []*Node{}
	queue = append(queue, n)

	nodes := []int{}

	for len(queue) > 0 {

		currentLevelSize := len(queue)

		for i, c := range queue {

			queue = queue[1:]

			if i == currentLevelSize-1 {
				nodes = append(nodes, c.Value)
			}

			if c.Left != nil {
				queue = append(queue, c.Left)
			}

			if c.Right != nil {
				queue = append(queue, c.Right)
			}

		}

	}

	fmt.Println("right most nodes", nodes)

}

func zigZagOrder(n *Node) {

	queue := []*Node{}
	queue = append(queue, n)

	result := map[int][]int{}

	level := 1

	for len(queue) > 0 {

		nextLevel := []*Node{}

		for _, c := range queue {

			if c.Left != nil {
				nextLevel = append(nextLevel, c.Left)
			}

			if c.Right != nil {
				nextLevel = append(nextLevel, c.Right)
			}

		}

		if level%2 == 0 {
			for i := len(queue) - 1; i >= 0; i-- {
				result[level] = append(result[level], queue[i].Value)
			}
		} else {
			for i := 0; i < len(queue); i++ {
				result[level] = append(result[level], queue[i].Value)
			}
		}

		level++

		queue = nextLevel

	}

	fmt.Println("level wise value", result)
}

func maxWidth(n *Node) {

	type childNodes struct {
		n     *Node
		index int
	}

	queue := []childNodes{}
	queue = append(queue, childNodes{n: n, index: 1})

	maxWidth := 0

	for len(queue) > 0 {

		start := queue[0].index
		end := queue[len(queue)-1].index
		width := end - start + 1
		if width > maxWidth {
			maxWidth = width
		}

		nextLevel := []childNodes{}

		for _, c := range queue {

			if c.n.Left != nil {
				nextLevel = append(nextLevel, childNodes{n: c.n.Left, index: c.index * 2})
			}

			if c.n.Right != nil {
				nextLevel = append(nextLevel, childNodes{n: c.n.Right, index: c.index*2 + 1})
			}

		}

		queue = nextLevel

	}

	fmt.Println("Max width is ", maxWidth)

}

func main() {

	n1 := Node{Value: 12}
	n2 := Node{Value: 5}
	n3 := Node{Value: 15}

	n1.Left = &n2
	n1.Right = &n3

	n4 := Node{Value: 3}
	n5 := Node{Value: 11}
	n6 := Node{Value: 13}

	n2.Left = &n4
	n2.Right = &n5

	n3.Left = &n6

	n3.Right = &Node{Value: 16}

	bfsTraversal(&n1)
	levelOrderSum(&n1)
	rightMostNode(&n1)
	zigZagOrder(&n1)
	maxWidth(&n1)

}
