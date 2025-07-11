package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3[a2[c]]"

	stack := []string{}

	var result string

	for _, d := range s {

		c := string(d)

		if c != "]" {
			stack = append(stack, c)
		} else {

			cs := ""
			for {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if pop != "[" 
			}

			fmt.Println("stack", stack)

			n, _ := strconv.Atoi(string(pop))
			for i := 0; i < n; i++ {
				cs += cs
			}

			if len(stack) != 0 {
				stack = append(stack, cs)
			} else {
				result += cs
			}

			fmt.Println("result", stack)

		}

	}

	fmt.Println("result string", result)
}
