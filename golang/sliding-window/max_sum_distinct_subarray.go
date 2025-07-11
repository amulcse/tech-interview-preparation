package main

import "fmt"

func main() {

	nums := []int{3, 2, 2, 3, 4, 6, 7, 7, -1}
	k := 4

	start := 0
	state := map[int]int{}
	currentSum := 0
	maxSum := 0

	for end := range nums {

		state[nums[end]]++
		currentSum += nums[end]

		if end-start+1 == k {

			if len(state) == k {
				maxSum = max(maxSum, currentSum)
			}

			currentSum -= nums[start]
			state[nums[start]]--

			if state[nums[start]] == 0 {
				delete(state, nums[start])
			}

			start++

		}

	}

	fmt.Println("max sum ", maxSum)

}
