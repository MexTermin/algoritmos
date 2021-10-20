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
func (T *Linked) Lenght() int {
	return T.len // 1
}

// O(n)
func (T *Linked) Get(index int) interface{} {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}
	var node *Node = T.head           //1
	for i := index - 1; i >= 0; i-- { // n
		node = node.next
	}
	return node // 1
}

// O(n)
func (T *Linked) Delete(index int) {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}
	var node *Node = T.head           //1
	for i := index - 1; i >= 0; i-- { // n
		node = node.next
	}
	if node.prev == nil { // 1
		T.head = node.next // 1
	} else {
		node.prev.next = node.next // 1
		if node.next != nil {
			node.next.prev = node.prev
		}
	}
	node = nil
	T.len-- // 1
}

// O(n)
func (T *Linked) Insert(index int, value interface{}) {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}
	var node *Node = T.head           //1
	for i := index - 1; i >= 0; i-- { // n
		node = node.next
	}

	element := &Node{ // 1
		value: value,     // 1
		next:  node,      // 1
		prev:  node.prev, // 1
	}

	if node.prev == nil {
		T.head = element // 1
	} else {
		node.prev.next = element // 1
		if node.next != nil {
			node.next.prev = element
		}
	}
	T.len++
}

// O(1)
func (T *Linked) Add(value interface{}) {
	element := &Node{ // 1
		value: value,  // 1
		next:  nil,    // 1
		prev:  T.tail, // 1
	}
	if T.head == nil {
		T.head = element
	}
	if T.tail != nil {
		T.tail.next = element // 1
		T.tail = element      // 1
	} else {
		T.tail = element
	}
	T.len++ // 1
}

//-------------Implementacion de la estructura de la lista, y enlace de esta con la interfaz

func NewLinkedList() *Linked {
	var ListaStruc *Linked = &Linked{
		head: nil,
		tail: nil,
		len:  0,
	}
	implementInterface(ListaStruc)
	return ListaStruc
}

//-------------------------------------------------------------------------------------------
