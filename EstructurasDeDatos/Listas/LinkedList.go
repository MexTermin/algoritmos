package EstructurasDeDatos

import (
	utils "algoritmos/EstructurasDeDatos/Utils"
	"fmt"
)

type Linked struct {
	head *utils.NodeDouble
	tail *utils.NodeDouble
	len  int
}

// O(1)
func (T *Linked) Lenght() int {
	return T.len // 1
}

// O(n)
func (T *Linked) Get(index int) interface{} {
	if index >= T.len || index < 0 { // 1
		panic("Indice fuera de rango") // 1
	}
	if index == 0 { // 1
		return T.head.Value // 1
	} else if index == T.len-1 { // 1
		return T.tail.Value // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.Next // 1
		}
		return node.Value // 1
	}
}

// O(n)
func (T *Linked) Delete(index int) {
	if index >= T.len || index < 0 { // 1
		panic("Indice fuera de rango") // 1
	}
	if index == 0 { // 1
		T.head = T.head.Next // 1
		T.head.Prev = nil    // 1
	} else if index == T.len-1 { // 1
		T.tail = T.tail.Prev // 1
		T.tail.Next = nil    // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.Next // 1
		}
		node.Prev.Next = node.Next // 1
		node.Next.Prev = node.Prev // 1
	}
	T.len-- // 1
}

// O(n)
func (T *Linked) Insert(index int, value interface{}) {
	if index >= T.len || index < 0 { // 1
		panic("Indice fuera de rango") // 1
	}

	element := &utils.NodeDouble{ // 1
		Value: value, // 1
		Next:  nil,   // 1
		Prev:  nil,   // 1
	}
	if index == 0 { // 1
		element.Next = T.head // 1
		T.head.Prev = element // 1
		T.head = element      // 1
	} else if index == T.len-1 { // 1
		element.Prev = T.tail // 1
		T.tail.Next = element // 1
		T.tail = element      // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.Next // 1
		}
		element.Next = node      // 1
		element.Prev = node.Prev // 1
		node.Prev.Next = element // 1
		node.Prev = element      // 1
	}
	T.len++ // 1
}

// O(1)
func (T *Linked) Add(value interface{}) {
	element := &utils.NodeDouble{ // 1
		Value: value, // 1
		Next:  nil,   // 1
		Prev:  nil,   // 1
	}
	if T.head == nil && T.tail == nil { // 1
		T.head = element // 1
		T.tail = element // 1
	} else { // 1
		element.Prev = T.tail // 1
		T.tail.Next = element // 1
		T.tail = element      // 1
	}
	T.len++ // 1
}

// O(n)
func (T *Linked) DisplayHead() {
	element := T.head    // 1
	for element != nil { // n
		if element.Next != nil { // 1
			fmt.Print(element.Value, " -> ") // 1
		} else { // 1
			fmt.Print(element.Value) // 1
		}
		element = element.Next // 1
	}
	fmt.Println("") // 1
}

// O(n)
func (T *Linked) DisplayTails() {
	element := T.tail    // 1
	for element != nil { // n
		if element.Prev != nil { // 1
			fmt.Print(element.Value, " <- ") // 1
		} else { // 1
			fmt.Print(element.Value) // 1
		}
		element = element.Prev // 1
	}
	fmt.Println("") // 1
}

//-------------Implementacion de la estructura de la lista, y enlace de esta con la interfaz
// O(1)
func NewLinkedList() *Linked {
	var ListaStruc *Linked = &Linked{ // 1
		head: nil, // 1
		tail: nil, // 1
		len:  0,   // 1
	}
	implementInterface(ListaStruc) // 1
	return ListaStruc              // 1
}

//-------------------------------------------------------------------------------------------
