package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SingleLL struct {
	head *Node
	tail *Node
}

// InsertAtEnd adds a new node at the end
func (ll *SingleLL) InsertAtEnd(data int) {
	newNode := &Node{data: data, next: nil} // Fix: Create a pointer directly

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	ll.tail.next = newNode
	ll.tail = newNode
}

// InsertAtStart adds a new node at the beginning
func (ll *SingleLL) InsertAtStart(data int) {
	newNode := &Node{data: data, next: ll.head} // Fix: Create a pointer directly

	ll.head = newNode
	if ll.tail == nil { // Fix: If list was empty, update tail
		ll.tail = newNode
	}
}

// Delete removes a node with a given value
func (ll *SingleLL) Delete(data int) {
	if ll.head == nil {
		return
	}

	// If head needs to be deleted
	if ll.head.data == data {
		ll.head = ll.head.next
		if ll.head == nil { // Fix: If list becomes empty, tail should also be nil
			ll.tail = nil
		}
		return
	}

	current := ll.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next != nil { // Node found
		if current.next == ll.tail { // Fix: If deleting the tail node, update tail
			ll.tail = current
		}
		current.next = current.next.next
	}
}

// Display prints the linked list
func (ll *SingleLL) Display() {
	fmt.Println("\nSingle Linked List:")
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
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

	ll.Delete(133) // No change, since 133 is not in the list
	ll.Delete(11)  // Deletes head
	ll.Delete(20)  // Deletes tail
	ll.InsertAtEnd(21)

	ll.Display()
}
