package EstructurasDeDatos

import utils "algoritmos/EstructurasDeDatos/Utils"

type LinkedQueue struct {
	head *utils.NodeSimple
	tail *utils.NodeSimple
}

// O(1)
func (Queue *LinkedQueue) IsEmpty() bool {
	return Queue.head != nil // 1
}

// O(1)
func (Queue *LinkedQueue) EnQueue(value interface{}) {
	var node = &utils.NodeSimple{ // 1
		Value: value, // 1
		Next:  nil,   // 1
	}
	if Queue.head == nil { // 1
		Queue.head = node // 1
		Queue.tail = node // 1
	} else {
		Queue.tail.Next = node // 1
		Queue.tail = node      // 1
	}
}

// O(1)
func (Queue *LinkedQueue) DeQueue() interface{} {
	if Queue.head == nil { // 1
		panic("La cola esta vacia") // 1
	}
	var result interface{} = Queue.head.Value // 1
	Queue.head = Queue.head.Next              // 1
	if Queue.head == nil {                    // 1
		Queue.tail = nil // 1
	}
	return result // 1
}

// O(1)
func (Queue *LinkedQueue) Head() interface{} {
	if Queue.head == nil { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.head.Value //1
}

// O(1)
func (Queue *LinkedQueue) Tail() interface{} {
	if Queue.tail == nil { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.tail.Value //1
}

//-------------Implementacion de la estructura de colas, y enlace de esta con la interfaz Queue
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