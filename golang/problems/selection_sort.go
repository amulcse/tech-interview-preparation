package main

import "fmt"

func main() {
	data := []int{5, 7, 5, 3, 2, 1}

	n := len(data)

	for i := 0; i < n-1; i++ {

		smallElementIndex := i

		for j := i + 1; j < n; j++ {

			if data[j] < data[smallElementIndex] {
				smallElementIndex = j
			}

		}

		if i != smallElementIndex {
			data[smallElementIndex], data[i] = data[i], data[smallElementIndex]
		}

	}

	fmt.Println("sorted data", data)
}
