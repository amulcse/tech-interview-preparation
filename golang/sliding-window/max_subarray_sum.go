package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}

	start := 0
	k := 3
	currentSum := 0
	maxSum := 0

	for end := range nums {

		currentSum += nums[end]

		if end >= k {
			currentSum -= nums[start]
			start++
		}

		maxSum = max(maxSum, currentSum)

	}

	fmt.Println("Maximum sum ", maxSum)
}
