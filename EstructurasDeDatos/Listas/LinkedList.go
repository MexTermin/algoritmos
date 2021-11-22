package Listas

import (
	utils "algoritmos/EstructurasDeDatos/Utils"
)

type Linked struct {
	head *utils.DoubleNode
	tail *utils.DoubleNode
	len  int
}

// O(1)
func (T *Linked) Lenght() int {
	return T.len // 1
}

// O(n)
func (T *Linked) Get(index int) interface{} {
	if index >= T.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}
	if index == 0 { // 1
		return T.head.GetValue() // 1
	} else if index == T.len - 1 { // 1
		return T.tail.GetValue() // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.GetNext() // 1
		}
		return node.GetValue() // 1
	}
}

// O(n)
func (T *Linked) Delete(index int) {
	if index >= T.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}
	if index == 0 { // 1
		T.head = T.head.GetNext() // 1
		T.head.SetPrev(nil)       // 1
	} else if index == T.len - 1 { // 1
		T.tail = T.tail.GetPrev() // 1
		T.tail.SetNext(nil)       // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.GetNext() // 1
		}
		node.GetPrev().SetNext(node.GetNext()) // 1
		node.GetNext().SetPrev(node.GetPrev()) // 1
	}
	T.len-- // 1
}

// O(n)
func (T *Linked) Insert(index int, value interface{}) {
	if index >= T.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}

	element := &utils.DoubleNode{} //1
	element.SetValue(value)        // 1
	if index == 0 {                // 1
		element.SetNext(T.head) // 1
		T.head.SetPrev(element) // 1
		T.head = element        // 1
	} else if index == T.len - 1 { // 1
		element.SetPrev(T.tail) // 1
		T.tail.SetNext(element) // 1
		T.tail = element        // 1
	} else { // 1
		node := T.head               //1
		for i := 0; i < index; i++ { // n
			node = node.GetNext() // 1
		}
		element.SetNext(node)           // 1
		element.SetPrev(node.GetPrev()) // 1
		node.GetPrev().SetNext(element) // 1
		node.SetPrev(element)           // 1
	}
	T.len++ // 1
}

// O(1)
func (T *Linked) Add(value interface{}) {
	element := &utils.DoubleNode{}      // 1
	element.SetValue(value)             // 1
	if T.head == nil && T.tail == nil { // 1
		T.head = element // 1
		T.tail = element // 1
	} else { // 1
		element.SetPrev(T.tail) // 1
		T.tail.SetNext(element) // 1
		T.tail = element        // 1
	}
	T.len++ // 1
}

//-------------Implementación de la estructura de la lista y enlace de esta con la interfaz
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
