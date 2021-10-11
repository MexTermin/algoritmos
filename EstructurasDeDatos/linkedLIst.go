package EstructurasDeDatos

import "fmt"

type node struct {
	prev  *node
	next  *node
	value interface{}
}

type Linked struct {
	head *node
	tail *node
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

func (T *Linked) Len() int {
	return T.len
}

// O(n)
func (T *Linked) Adiccion(value interface{}) {
	element := &node{ // 1
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
func (T *Linked) FindNode(value int) *node {
	head := T.head    //1
	for head != nil { // n
		if head.value == value { // n
			return head // n
		}
		head = head.next // n
	}
	return nil // 1
}

