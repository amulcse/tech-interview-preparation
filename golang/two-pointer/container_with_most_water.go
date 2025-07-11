package main

import "fmt"

func main() {

	data := []int{1, 4, 2, 5, 7}

	maxWaterLevel := 0

	left := 0
	right := len(data) - 1

	for left < right {

		width := right - left
		height := min(data[left], data[right])
		area := width * height
		maxWaterLevel = max(area, maxWaterLevel)

		if data[left] < data[right] {
			left++
		} else {
			right--
		}

	}

	fmt.Println("Maximum water level could be ", maxWaterLevel)

}
