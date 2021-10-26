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

// O(1)
func (T *Linked) Lenght() int {
	return T.len // 1
}

// O(n)
func (T *Linked) Get(index int) interface{} {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}
	if index == 0 { // 1
		return T.head.value // 1
	} else if index == T.len-1 { // 1
		return T.tail.value // 1
	} else {
		var node *Node = T.head      //1
		for i := 0; i < index; i++ { // n
			node = node.next // 1
		}
		return node.value // 1
	}
}

// O(n)
func (T *Linked) Delete(index int) {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}
	if index == 0 { // 1
		T.head = T.head.next // 1
		T.head.prev = nil    // 1
	} else if index == T.len-1 { // 1
		T.tail = T.tail.prev // 1
		T.tail.next = nil    // 1
	} else {
		var node *Node = T.head      //1
		for i := 0; i < index; i++ { // n
			node = node.next // 1
		}
		node.prev.next = node.next // 1
		node.next.prev = node.prev // 1
	}
	T.len-- // 1
}

// O(n)
func (T *Linked) Insert(index int, value interface{}) {
	if index >= T.len || index < 0 {
		panic("Indice fuera de rango")
	}

	element := &Node{ // 1
		value: value, // 1
		next:  nil,   // 1
		prev:  nil,   // 1
	}
	if index == 0 {
		element.next = T.head // 1
		T.head.prev = element // 1
		T.head = element      // 1
	} else if index == T.len-1 {
		element.prev = T.tail // 1
		T.tail.next = element // 1
		T.tail = element      // 1
	} else {
		var node *Node = T.head      //1
		for i := 0; i < index; i++ { // n
			node = node.next // 1
		}
		element.next = node      // 1
		element.prev = node.prev // 1
		node.prev.next = element // 1
		node.prev = element      // 1
	}
	T.len++ // 1
}

// O(1)
func (T *Linked) Add(value interface{}) {
	element := &Node{ // 1
		value: value, // 1
		next:  nil,   // 1
		prev:  nil,   // 1
	}
	if T.head == nil && T.tail == nil {
		T.head = element // 1
		T.tail = element // 1
	} else {
		element.prev = T.tail // 1
		T.tail.next = element // 1
		T.tail = element      // 1
	}
	T.len++ // 1
}

func (T *Linked) DisplayHead() {
	element := T.head

	for element != nil { // n
		if element.next != nil {
			fmt.Print(element.value, " -> ")
		} else {
			fmt.Print(element.value)
		}
		element = element.next
	}
	fmt.Println("")
}

func (T *Linked) DisplayTails() {
	element := T.tail

	for element != nil { // n
		if element.prev != nil {
			fmt.Print(element.value, " <- ")
		} else {
			fmt.Print(element.value)
		}
		element = element.prev
	}
	fmt.Println("")
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
