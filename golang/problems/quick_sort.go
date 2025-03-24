package main

import "fmt"

func partition(data []int, low, high int) int {
	pivot := data[high]

	i := low - 1

	for j := low; j <= high; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}

	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}

func quicksort(data []int, low, high int) {

	if low < high {
		p := partition(data, low, high)
		quicksort(data, low, p-1)
		quicksort(data, p+1, high)
	}

}

func main() {

	data := []int{5, 1, 2, 9, 7}
	quicksort(data, 0, len(data)-1)
	fmt.Println(data)

}
