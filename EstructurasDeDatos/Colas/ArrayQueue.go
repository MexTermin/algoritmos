package Colas

import utils "algoritmos/EstructurasDeDatos/Utils"

type Queue struct {
	children     []interface{}
	lastPosition int
}

// O(1)
func (Queue *Queue) IsEmpty() bool {
	return Queue.lastPosition == 0 // 1
}

// O(n)
func (Queue *Queue) EnQueue(value interface{}) {
	if Queue.lastPosition == len(Queue.children) { // 1
		Queue.children = utils.ResizeArray(Queue.children, len(Queue.children) * 2, Queue.lastPosition - 1) // n
	}
	Queue.children[Queue.lastPosition] = value // 1
	Queue.lastPosition++                       // 1
}

// O(n)
func (Queue *Queue) DeQueue() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola está vacía") // 1
	}
	result := Queue.children[0]                             // 1
	for index := 0; index < Queue.lastPosition - 1; index++ { // n
		Queue.children[index] = Queue.children[index + 1] // 1
	}
	if len(Queue.children) / 2 > Queue.lastPosition && Queue.lastPosition > 10 { // 1
		Queue.children = utils.ResizeArray(Queue.children, Queue.lastPosition * 2, Queue.lastPosition - 1) // n
	}
	Queue.children[Queue.lastPosition - 1] = nil // 1
	Queue.lastPosition--                       // 1
	return result                              // 1
}

// O(1)
func (Queue *Queue) Head() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola está vacía") // 1
	}
	return Queue.children[0] // 1
}

// O(1)
func (Queue *Queue) Tail() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola está vacía") // 1
	}
	return Queue.children[Queue.lastPosition - 1] // 1
}

//-------------Instánciador del Queue y enlace de esta con la interfaz IQueue
// O(1)
func NewArrayQueue() *Queue {
	var Queue *Queue = &Queue{ // 1
		children:     make([]interface{}, 10), // 1
		lastPosition: 0,
	}
	implementIQueue(Queue) // 1
	return Queue           // 1
}

//-------------------------------------------------------------------------------------------
