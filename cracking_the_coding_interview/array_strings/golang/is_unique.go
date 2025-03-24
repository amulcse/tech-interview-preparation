/*
Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/
package main

import "fmt"

func isUnique(text string) bool {

	if len(text) > 26 {
		return false
	}

	characterCount := make(map[string]int)

	for _, c := range text {
		char := string(c)
		if _, found := characterCount[char]; !found {
			characterCount[char] = 0
		}
		characterCount[char]++

		if characterCount[char] > 1 {
			return false
		}
	}

	return true

}

func main() {

	fmt.Println(isUnique("aA"))

}
