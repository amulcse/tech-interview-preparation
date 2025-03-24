package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}

	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return result
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Top() int {

	if len(s.items) == 0 {
		return -1
	}

	return s.items[0]
}

func main() {

	stack := Stack{}

	stack.Push(10)
	stack.Push(13)
	stack.Push(15)
	stack.Push(1)
	stack.Push(2)

	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Pop:", stack.Pop())

	stack.Push(5)

	fmt.Printf("Stack size: %d", stack.Size())

}
