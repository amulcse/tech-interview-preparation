// Given a list of sorted movies from multiple providers based on rating,
// produce a list of movies that are sorted across all providers.

package main

import "fmt"

type movie struct {
	name    string
	ratting int
}

func main() {

	result := []movie{}

	prime := []movie{
		{name: "moview 1", ratting: 5},
		{name: "moview 2", ratting: 4},
	}

	netflix := []movie{
		{name: "moview 3", ratting: 4},
		{name: "moview 5", ratting: 3},
	}

	totalRecords := len(prime) + len(netflix)

	providers := map[string][]movie{"prime": prime, "netflix": netflix}

	indexTracker := map[string]int{"prime": 0, "netflix": 0}

	for i := 0; i < totalRecords; i++ {

		currentMovieIndex := -1
		currentProvider := ""

		for p, movies := range providers {

			currentIndex := indexTracker[p]

			if currentIndex >= len(movies) {
				continue
			}

			if currentMovieIndex == -1 {
				currentMovieIndex = currentIndex
				currentProvider = p
			}

			if providers[p][currentIndex].ratting >= providers[currentProvider][currentMovieIndex].ratting {
				currentMovieIndex = currentIndex
				currentProvider = p
			}

		}

		result = append(result, providers[currentProvider][currentMovieIndex])
		indexTracker[currentProvider]++

	}

	fmt.Println("Results", result)

}
