package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func sumOfNodes(n *Node) int {

	if n == nil {
		return 0
	}

	return sumOfNodes(n.Left) + sumOfNodes(n.Right) + n.Value

}

func maxValue(n *Node) int {

	if n == nil {
		return math.MinInt
	}

	if n.Left == nil && n.Right == nil {
		return n.Value
	}

	return max(maxValue(n.Left), maxValue(n.Right), n.Value)

}

func targetTotalPathExist(n *Node, target int) bool {

	if n == nil {
		return false
	}

	target -= n.Value

	if n.Left == nil && n.Right == nil && target == 0 {
		return true
	}

	fmt.Println("the sum is,", target)

	left := targetTotalPathExist(n.Left, target)
	right := targetTotalPathExist(n.Right, target)

	if left || right {
		return true
	}

	return false

}

func maxDepth(n *Node) int {

	if n == nil {
		return 0
	}

	return max(maxDepth(n.Left), maxDepth(n.Right)) + 1

}

func noOfGoodNodes(n *Node, max int) int {

	if n == nil {
		return 0
	}

	count := 0
	if n.Value > max {
		max = n.Value
		count = 1
	}

	return noOfGoodNodes(n.Left, max) + noOfGoodNodes(n.Right, max) + count

}

func isValidBST(n *Node, min int, max int) bool {

	if n == nil {
		return true
	}

	if !(n.Value < max && n.Value > min) {
		return false
	}

	return isValidBST(n.Left, min, n.Value) && isValidBST(n.Right, n.Value, max)

}

func sumOfTilt(n *Node) int {

	var sum int

	var dfs func(n *Node) int

	dfs = func(n *Node) int {

		if n == nil {
			return 0
		}

		left := dfs(n.Left)
		right := dfs(n.Right)
		sum += int(math.Abs(float64(right - left)))

		return left + right + n.Value

	}

	dfs(n)

	return sum

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

	target := 21

	fmt.Println("Sum of node is :", sumOfNodes(&n1))
	fmt.Println("Max value is :", maxValue(&n1))
	fmt.Println("Max depth is :", maxDepth(&n1))
	fmt.Println("Target path exist :", targetTotalPathExist(&n1, target))
	fmt.Println("Number of good nodes :", noOfGoodNodes(&n1, math.MinInt))
	fmt.Println("Is valid BST :", isValidBST(&n1, math.MinInt, math.MaxInt))
	fmt.Println("Sum of tilt :", sumOfTilt(&n1))

	fmt.Println("min int", math.MinInt)

}
