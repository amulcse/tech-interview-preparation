package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int               { return len(h) }
func (h IntHeap) Less(i int, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	h := &IntHeap{5, 7, 4, 2, 4, 5}
	heap.Init(h)

	heap.Push(h, 23)
	heap.Push(h, 29)
	heap.Push(h, 22)
	heap.Push(h, 24)
	heap.Push(h, -3)

	for h.Len() > 0 {
		fmt.Println("heap pop", heap.Pop(h))
	}

}
