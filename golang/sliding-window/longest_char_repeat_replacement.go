package main

import "fmt"

func main() {

	s := "ABABABABBAA"
	k := 2

	start, end := 0, 0
	maxLength := 0

	for end <= len(s)-1 {

		if s[end] == s[start] {
			maxLength = max(maxLength, end-start+1)
			end++
		} else {
			if k > 0 {
				k--
				maxLength = max(maxLength, end-start+1)
				end++
			} else {
				c := s[start]
				for s[start] == c {
					start++
				}
				end = start
				k = 2
			}
		}

	}

	fmt.Println("maxLength", maxLength)
}
