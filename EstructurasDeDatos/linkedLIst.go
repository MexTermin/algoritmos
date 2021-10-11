package EstructurasDeDatos

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type Linked struct {
	head *Node
	tail *Node
	len  int
}

func (T *Linked) Display() {
	element := T.head

	for element != nil {
		if element.next != nil {
			fmt.Print(element.value, " -> ")
		} else {
			fmt.Print(element.value)
		}
		element = element.next
	}
	fmt.Println("")
}

// O(1)
func (T *Linked) Len() int {
	return T.len // 1
}

// O(n)
func (T *Linked) Push(value interface{}) {
	element := &Node{ // 1
		next:  T.head, // 1
		value: value,  // 1
	}
	if T.head != nil { // 1
		T.head.prev = element // 1
	}
	T.head = element // 1

	tails := T.head         // 1
	for tails.next != nil { // n
		tails = tails.next // n
	}
	T.len++        // 1
	T.tail = tails // 1
}

// O(n)
func (T *Linked) FindNode(value int) *Node {
	head := T.head    //1
	for head != nil { // n
		if head.value == value { // n
			return head // n
		}
		head = head.next // n
	}
	return nil // 1
}

// O(1)
func (T *Linked) InsertAfter(N *Node, value int) {
	element := &Node{ // 1
		value: value,  // 1
		next:  N.next, // 1
		prev:  N,      // 1
	}
	N.next = element // 1
	T.len++
}

// O(1)
func (T *Linked) InsertBefore(N *Node, value int) {
	element := &Node{ // 1
		value: value,  // 1
		next:  N,      // 1
		prev:  N.prev, // 1
	}
	if N.prev == nil {
		T.head = element // 1
	} else {
		N.prev.next = element // 1
	}
	T.len++
}
