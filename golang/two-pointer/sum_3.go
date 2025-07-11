package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	nums := []int{-2, 4, 1, 0, -1, 1, 2, -3, 0}
	slices.Sort(nums)

	result := [][]int{}
	founds := map[string]bool{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate first elements
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				triplet := []int{nums[i], nums[left], nums[right]}
				key := strconv.Itoa(triplet[0]) + "_" + strconv.Itoa(triplet[1]) + "_" + strconv.Itoa(triplet[2])

				if !founds[key] {
					founds[key] = true
					result = append(result, triplet)
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	fmt.Println("result is", result)
}
