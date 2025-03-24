package main

import (
	"fmt"
	"sort"
)

func solution1(a string, b string) bool {

	if len(a) != len(b) {
		return false
	}

	aSlice := []rune(a)
	sort.Slice(aSlice, func(a, b int) bool {
		return aSlice[a] < aSlice[b]
	})
	sortedA := string(aSlice)

	bSlice := []rune(b)
	sort.Slice(bSlice, func(a, b int) bool {
		return bSlice[a] < bSlice[b]
	})
	sortedB := string(bSlice)

	if sortedA != sortedB {
		return false
	}

	return true

}

func solution2(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aCharCountHashMap := make(map[string]int)
	bCharCountHashMap := make(map[string]int)
	n := len(a)

	for i := range n {
		if _, found := aCharCountHashMap[string(a[i])]; !found {
			aCharCountHashMap[string(a[i])] = 0
		}
		if _, found := bCharCountHashMap[string(b[i])]; !found {
			bCharCountHashMap[string(b[i])] = 0
		}
		aCharCountHashMap[string(a[i])]++
		bCharCountHashMap[string(a[i])]++
	}

	if len(aCharCountHashMap) != len(bCharCountHashMap) {
		return false
	}

	return true
}

func main() {

	string1 := "acaa"
	string2 := "aaac"
	fmt.Printf("%s is permutation of %s: %t", string1, string2, solution1(string1, string2))

}
