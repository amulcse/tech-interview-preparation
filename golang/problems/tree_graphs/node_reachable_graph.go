package main

import "fmt"

type Graph struct {
	edges map[string][]string
}

// BFS to check if the destination is reachable from source
func isReachable(graph Graph, source string, destination string) bool {

	if source == destination {
		return true
	}

	visited := make(map[string]bool)
	queue := []string{source}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range graph.edges[current] {
			if neighbor == destination {
				return true
			}

			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}

		}

	}

	return false

}

func main() {

	graph := Graph{
		edges: make(map[string][]string),
	}
	graph.edges["A"] = []string{"B", "C"}
	graph.edges["B"] = []string{"C"}
	graph.edges["D"] = []string{"D"}

	source := "A"
	destination := "D"

	fmt.Printf("Is there a path between %s & %s = %t\n", source, destination, isReachable(graph, source, destination))

}
