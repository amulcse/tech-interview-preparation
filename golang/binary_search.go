package main

import "fmt"

func binarySearch(data []int, value int) bool {

	left := 0
	right := len(data) - 1

	for left <= right {
		middle := (left + right) / 2
		if data[middle] == value {
			return true
		} else if data[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false

}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(data, 11)
	fmt.Println("result is", result)

}
