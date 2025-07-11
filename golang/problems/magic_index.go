package main

import "fmt"

func findMagicIndex(data []int) int {
	for i, v := range data {

		if i == v {
			return i
		}
	}

	return -1
}

func findMagicIndexBS(data []int, start int, end int) int {

	if start < 0 || end >= len(data) || start > end {
		return -1
	}

	mid := (start + end) / 2
	if data[mid] == mid {
		return mid
	}

	if data[mid] <= mid {
		return findMagicIndexBS(data, mid+1, end)
	} else {
		return findMagicIndexBS(data, start, end-1)
	}

}

func main() {

	data := []int{-5, -3, 0, 1, 3, 5}

	fmt.Println("Magic index is", findMagicIndexBS(data, 0, len(data)-1))

}
