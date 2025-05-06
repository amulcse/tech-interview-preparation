package main

import "fmt"

func getBuildOrder(projects []string, dependencies [][2]string) (error, []string) {

	graph := make(map[string][]string)
	inOrder := make(map[string]int)
	var buildOrder []string

	for _, project := range projects {
		graph[project] = []string{}
		inOrder[project] = 0
	}

	for _, dependency := range dependencies {
		parent, child := dependency[0], dependency[1]
		graph[parent] = append(graph[parent], child)
		inOrder[child]++
	}

	queue := []string{}
	for project, count := range inOrder {
		if count == 0 {
			queue = append(queue, project)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		buildOrder = append(buildOrder, current)

		for _, neighbor := range graph[current] {
			inOrder[neighbor]--

			if inOrder[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(buildOrder) != len(projects) {
		return fmt.Errorf("There is cycle"), []string{}
	}

	return nil, buildOrder

}

func main() {

	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][2]string{
		{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"},
	}

	err, result := getBuildOrder(projects, dependencies)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(result)

}
