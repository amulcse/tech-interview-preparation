package main

import "fmt"

func main() {
	temps := []int{65, 70, 68, 60, 55, 75, 80, 74}

	stack := []int{0}

	result := []int{}
	for range temps {
		result = append(result, 0)
	}

	for i, t := range temps {

		for len(stack) > 0 {
			pick := stack[len(stack)-1]
			if temps[pick] > t {
				break
			}
			result[pick] = i - pick
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

	}

	fmt.Println(result)
}
