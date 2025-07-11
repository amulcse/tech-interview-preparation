package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value     int
	index     int
	listIndex int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (pq *ItemHeap) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *ItemHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	a := []int{3, 4, 5}
	b := []int{2, 3, 5}
	c := []int{-1, 6}
	lists := [][]int{a, b, c}

	result := []int{}
	h := &ItemHeap{}
	heap.Init(h)

	for i, l := range lists {
		if len(l) > 0 {
			heap.Push(h, &Item{index: 0, value: l[0], listIndex: i})
		}
	}

	for h.Len() > 0 {
		minItem := heap.Pop(h).(*Item)
		result = append(result, minItem.value)

		if minItem.index+1 < len(lists[minItem.listIndex]) {
			newItem := &Item{
				index:     minItem.index + 1,
				value:     lists[minItem.listIndex][minItem.index+1],
				listIndex: minItem.listIndex,
			}
			heap.Push(h, newItem)
		}
	}

	fmt.Println(result)
}
