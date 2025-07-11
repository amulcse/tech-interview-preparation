package main

import "fmt"

func generateNeighborPattern(words []string) map[string]map[string]bool {

	patterns := map[string][]string{}

	for _, word := range words {

		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			patterns[pattern] = append(patterns[pattern], word)
		}

	}

	fmt.Println("patterns", patterns)

	neighbor := map[string]map[string]bool{}
	for _, word := range words {
		neighbor[word] = map[string]bool{}
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			if wordList, found := patterns[pattern]; found {
				for _, w := range wordList {
					if w != word {
						neighbor[word][w] = true
					}
				}
			}
		}
	}

	fmt.Println("patterns", neighbor)

	return neighbor

}

func findMinWordLadder(neighbors map[string]map[string]bool, begin string, end string) {

	queue := []string{}
	queue = append(queue, begin)

	steps := 0
	visited := map[string]bool{}
	found := false

	for len(queue) > 0 {

		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {

			word := queue[0]
			if word == end {
				found = true
				break
			}
			visited[word] = true
			queue = queue[1:]

			for n, _ := range neighbors[word] {
				if !visited[n] {
					queue = append(queue, n)
				}
			}

		}

		steps++

		if found {
			break
		}

	}

	if found {
		fmt.Printf("The word %s can reach word %s by steps %d", begin, end, steps)
	} else {
		fmt.Printf("The word %s can't reach word %s", begin, end)
	}

}

func main() {

	words := []string{"hit", "hot", "dot", "dog", "lot", "log", "cog"}
	begin := "hit"
	end := "cog"

	neighbors := generateNeighborPattern(words)
	findMinWordLadder(neighbors, begin, end)

}
