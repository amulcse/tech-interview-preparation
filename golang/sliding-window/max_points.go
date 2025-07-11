package main

import "fmt"

func main() {

	nums := []int{2, 11, 4, 5, 3, 9, 2}
	k := 3

	total := 0

	for i := range nums {
		total += nums[i]
	}

	start := 0
	state := 0
	maxPoints := 0

	for end := range nums {

		state += nums[end]

		if end-start+1 == len(nums)-k {

			maxPoints = max(maxPoints, total-state)
			state -= nums[start]
			start++

		}

	}

	fmt.Println("max points", maxPoints)

}
