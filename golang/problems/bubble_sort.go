package main

import "fmt"

func main() {
	data := []int{1, 3, 9, 7, 13}

	n := len(data)
	for i := 0; i < n-1; i++ {

		swapped := false
		for j := 0; j < n-i-1; j++ {

			if data[j] >= data[j+1] {
				temp := data[j+1]
				data[j+1] = data[j]
				data[j] = temp
				swapped = true
			}

		}
		// if two elements not swapped by inner loop then break
		if !swapped {
			break
		}

	}

	fmt.Println("sorted data", data)
}
