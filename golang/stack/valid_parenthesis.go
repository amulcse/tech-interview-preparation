package main

import "fmt"

func main() {
	s := "))"

	ls := 0

	stack := []int{-1}

	for i, r := range s {
		if r == '(' {
			stack = append(stack, i)
		} else {

			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ls = max(ls, i-stack[len(stack)-1])
			}

		}
	}

	fmt.Println("ls", ls)
}
