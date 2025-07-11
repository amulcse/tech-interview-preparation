package main

import "fmt"

func main() {
	s := "abddcbsyd"

	substringMap := map[byte]bool{}

	start := 0
	maxLength := 0

	for end := range s {

		c := s[end]

		// if duplicate found, remove all the character until this point
		for substringMap[c] {
			delete(substringMap, s[start])
			start++
		}

		substringMap[c] = true
		maxLength = max(maxLength, end-start+1)

	}

	fmt.Println("Max length is -- ", maxLength)

}
