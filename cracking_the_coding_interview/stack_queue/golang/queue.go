package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) EnQueue(data int) {
	n := &Node{data: data}

	if q.front == nil {
		q.front = n
	}

	if q.rear == nil {
		q.rear = n
		return
	}

	q.rear.next = n
	q.rear = n
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) DeQueue() int {
	if q.front == nil {
		return -1
	}

	data := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return data
}

func main() {

	queue := Queue{}
	queue.EnQueue(10)
	queue.EnQueue(10)
	queue.EnQueue(13)
	queue.EnQueue(10)

	fmt.Println("Dequeue:", queue.DeQueue())
	fmt.Println("Dequeue:", queue.DeQueue())
	fmt.Println("Dequeue:", queue.DeQueue())

}
