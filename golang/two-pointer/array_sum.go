package main

import "fmt"

func main() {

	data := []int{1, 2, 3, 5, 7, 9, 13}
	target := 25

	found := false

	left := 0
	right := len(data) - 1

	for left < right {

		currentSum := data[left] + data[right]

		if currentSum == target {
			found = true
			break
		}

		if currentSum < target {
			left++
		} else {
			right++
		}

	}

	fmt.Println("Target found:", found)

}
