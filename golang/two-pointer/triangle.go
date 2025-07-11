package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{11, 4, 9, 6, 15, 18}

	sort.Ints(nums)
	count := 0

	for i := range nums {

		left := i + 1
		right := len(nums) - 1

		for left < right {

			if nums[left]+nums[right] > nums[i] {
				count = count + right - left
				right--
			} else {
				left++
			}

		}

	}
	fmt.Println("count ", count)

}
