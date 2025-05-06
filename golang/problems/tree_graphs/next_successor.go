// Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a
// binary search tree. You may assume that each node has a link to its parent.
package main

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BinarySearchTree struct {
	Root *Node
}

func findInOrderSuccessor(sourceNode *Node, n *Node) *Node {

	if n == nil {
		return nil
	}

	rightChild := n.Right
	parentNode := n.Parent
	if parentNode != nil && parentNode.Value > sourceNode.Value {
		return parentNode
	}

	if rightChild != nil && rightChild.Value > sourceNode.Value {
		return rightChild
	}

	return findInOrderSuccessor(sourceNode, n.Parent)

}

func main() {

	bst := BinarySearchTree{}

	bst.Root = &Node{Value: 15}
	bst.Root.Left = &Node{Value: 10, Parent: bst.Root}
	bst.Root.Right = &Node{Value: 25, Parent: bst.Root}

	bst.Root.Left.Left = &Node{Value: 5, Parent: bst.Root.Left}
	bst.Root.Left.Right = &Node{Value: 13, Parent: bst.Root.Left}

	bst.Root.Right.Left = &Node{Value: 23, Parent: bst.Root.Right}
	// bst.Root.Right.Right = &Node{Value: 27, Parent: bst.Root.Right}

	successor := findInOrderSuccessor(bst.Root.Right, bst.Root.Right)
	if successor != nil {
		fmt.Println("successor value == ", successor.Value)
	} else {
		fmt.Println("No successor")
	}

}
