package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SingleLL struct {
	head *Node
}

func (ll *SingleLL) InsertAtEnd(data int) {
	newNode := Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = &newNode
}

func (ll *SingleLL) InsertAtStart(data int) {
	newNode := Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	newNode.next = ll.head
	ll.head = &newNode
}

func (ll *SingleLL) Delete(data int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (ll *SingleLL) Display() {
	fmt.Println("\nSingle Linked List")
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
}

func main() {

	ll := SingleLL{}
	ll.InsertAtEnd(12)
	ll.InsertAtEnd(13)
	ll.InsertAtEnd(15)
	ll.InsertAtStart(15)
	ll.InsertAtStart(11)
	ll.InsertAtEnd(20)

	ll.Display()

	ll.Delete(133)
	ll.Delete(11)
	ll.Display()

}
