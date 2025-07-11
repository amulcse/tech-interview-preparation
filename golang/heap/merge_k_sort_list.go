package main

import (
	"container/heap"
	"fmt"
)

type Movie struct {
	name   string
	rating int
}

type Item struct {
	movie      Movie
	listIndex  int
	movieIndex int
}

type ItemHeap []Item

func (h ItemHeap) Len() int { return len(h) }
func (h ItemHeap) Less(i, j int) bool {
	return h[i].movie.rating > h[j].movie.rating // max-heap
}
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {

	Netflix := []Movie{
		{name: "FF 6", rating: 5},
		{name: "FF 7", rating: 4},
		{name: "FF 10", rating: 1},
	}

	Prime := []Movie{
		{name: "MI 4", rating: 5},
		{name: "MI 3", rating: 2},
		{name: "FF 10", rating: 1},
	}

	lists := [][]Movie{Netflix, Prime}

	h := &ItemHeap{}
	heap.Init(h)

	// Push first movie from each list
	for i, list := range lists {
		if len(list) > 0 {
			heap.Push(h, Item{movie: list[0], listIndex: i, movieIndex: 0})
		}

		fmt.Println("heap length", h.Len())
	}

	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		fmt.Println("movie:", item.movie)

		// Push next movie from the same list
		if item.movieIndex+1 < len(lists[item.listIndex]) {
			nextIndex := item.movieIndex + 1
			heap.Push(h, Item{
				movie:      lists[item.listIndex][nextIndex],
				listIndex:  item.listIndex,
				movieIndex: nextIndex,
			})
		}
	}

}
