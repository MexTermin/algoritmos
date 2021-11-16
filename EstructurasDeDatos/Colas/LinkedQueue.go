package EstructurasDeDatos

import utils "algoritmos/EstructurasDeDatos/Utils"

type LinkedQueue struct {
	head *utils.SimpleNode
	tail *utils.SimpleNode
}

// O(1)
func (Queue *LinkedQueue) IsEmpty() bool {
	return Queue.head == nil // 1
}

// O(1)
func (Queue *LinkedQueue) EnQueue(value interface{}) {
	var node = &utils.SimpleNode{} // 1
	node.SetValue(value)           // 1
	if Queue.head == nil {         // 1
		Queue.head = node // 1
		Queue.tail = node // 1
	} else { // 1
		Queue.tail.SetNext(node) // 1
		Queue.tail = node        // 1
	}
}

// O(1)
func (Queue *LinkedQueue) DeQueue() interface{} {
	if Queue.head == nil { // 1
		panic("La cola está vacía") // 1
	}
	var result interface{} = Queue.head.GetValue() // 1
	Queue.head = Queue.head.GetNext()              // 1
	if Queue.head == nil {                         // 1
		Queue.tail = nil // 1
	}
	return result // 1
}

// O(1)
func (Queue *LinkedQueue) Head() interface{} {
	if Queue.head == nil { // 1
		panic("La cola está vacía") // 1
	}
	return Queue.head.GetValue() //1
}

// O(1)
func (Queue *LinkedQueue) Tail() interface{} {
	if Queue.tail == nil { // 1
		panic("La cola está vacía") // 1
	}
	return Queue.tail.GetValue() //1
}

//-------------Implementación de la estructura de colas y enlace de esta con la interfaz Queue
// O(1)
func NewLinkedQueue() *LinkedQueue {
	var QueueStruc *LinkedQueue = &LinkedQueue{ // 1
		head: nil, // 1
		tail: nil, // 1
	}
	implementIQueue(QueueStruc) // 1
	return QueueStruc           // 1
}
//-------------------------------------------------------------------------------------------
