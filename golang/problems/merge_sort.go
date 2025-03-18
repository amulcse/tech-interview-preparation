package main

import "fmt"

// Merge two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0

	// Merge the two slices while there are elements in both
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements in left or right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Merge Sort function that recursively splits and merges
func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	// Find the middle index to split the array
	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	// Recursively sort the left and right parts
	left = mergeSort(left)
	right = mergeSort(right)

	// Merge the sorted parts and return
	return merge(left, right)
}

func main() {
	data := []int{3, 1, 5, 7, 2}

	// Perform merge sort
	sortData := mergeSort(data)

	// Print the sorted result
	fmt.Println("Sorted data:", sortData)
}
