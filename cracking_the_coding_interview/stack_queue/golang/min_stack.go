package main

import (
	"fmt"
)

// MinStack struct
type MinStack struct {
	stack    []int // Main stack
	minStack []int // Stack to track min values
}

// Constructor initializes a new MinStack
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push adds an element to the stack
func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	// Push to minStack if it's empty or val is <= current min
	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

// Pop removes the top element from the stack
func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	// Remove from minStack if the popped element was the min
	if ms.stack[len(ms.stack)-1] == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}

	// Remove from main stack
	ms.stack = ms.stack[:len(ms.stack)-1]
}

// Min returns the minimum element in the stack
func (ms *MinStack) Min() int {
	if len(ms.minStack) == 0 {
		panic("Stack is empty")
	}
	return ms.minStack[len(ms.minStack)-1]
}

// Top returns the top element of the stack
func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		panic("Stack is empty")
	}
	return ms.stack[len(ms.stack)-1]
}

// Main function to test MinStack
func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println("Min:", s.Min()) // Output: 1
	s.Pop()
	fmt.Println("Min:", s.Min()) // Output: 2
}
