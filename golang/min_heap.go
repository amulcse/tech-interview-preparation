package main

import (
	"fmt"
)

type MinHeap struct {
	nodes []int
}

func (h *MinHeap) MinValue() int {
	if len(h.nodes) == 0 {
		return -1
	}

	return h.nodes[0]
}

func (h *MinHeap) Insert(data int) {

	h.nodes = append(h.nodes, data)
	h.BubbleUp(len(h.nodes) - 1)

}

func (h *MinHeap) Extract() int {

	if len(h.nodes) == 0 {
		return -1
	}

	min := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.BubbleDown(0)
	return min

}

func (h *MinHeap) BubbleUp(index int) {

	for index > 0 {
		parent := (index - 1) / 2

		if h.nodes[index] > h.nodes[parent] {
			break
		}

		h.nodes[index], h.nodes[parent] = h.nodes[parent], h.nodes[index]
		index = parent

	}

}

func (h *MinHeap) BubbleDown(index int) {

	size := len(h.nodes)

	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < size && h.nodes[left] < h.nodes[smallest] {
			smallest = left
		}

		if right < size && h.nodes[right] < h.nodes[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.nodes[index], h.nodes[smallest] = h.nodes[smallest], h.nodes[index]
		index = smallest

	}

}

func main() {

	h := MinHeap{}
	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(1)
	h.Extract()
	h.Extract()
	h.Extract()

	fmt.Println("Min value:", h.MinValue())

}
