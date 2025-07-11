package main

import "fmt"

var sum map[int]int

func calculateNumberOfWays(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	if _, exit := sum[n]; exit {
		return sum[n]
	}

	sum[n] = calculateNumberOfWays(n-1) + calculateNumberOfWays(n-2) + calculateNumberOfWays(n-3)

	return sum[n]

}

func main() {

	sum = make(map[int]int)
	fmt.Println(calculateNumberOfWays(22220000000000))
}
