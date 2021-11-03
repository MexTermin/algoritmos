package EstructurasDeDatos

import utils "algoritmos/EstructurasDeDatos/Utils"

type ArrayQueue struct {
	children     []interface{}
	lastPosition int
}

// O(1)
func (Queue *ArrayQueue) IsEmpty() bool {
	return Queue.lastPosition == 0 // 1
}

// O(n)
func (Queue *ArrayQueue) EnQueue(value interface{}) {
	if Queue.lastPosition == len(Queue.children) { // 1
		Queue.children = utils.ResizeArray(Queue.children, len(Queue.children)*2) // n
	}
	Queue.children[Queue.lastPosition] = value // 1
	Queue.lastPosition++                       // 1
}

// O(n)
func (Queue *ArrayQueue) DeQueue() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola esta vacia") // 1
	}
	result := Queue.children[0]        // 1
	index := 0                         // 1
	for Queue.children[index] != nil { // n
		Queue.children[index] = Queue.children[index+1] // 1
		index++                                         // 1
	}
	if Queue.children[(len(Queue.children)/2)+1] != nil { // 1
		Queue.children = utils.ResizeArray(Queue.children, len(Queue.children)*2) // n
	}
	Queue.lastPosition-- // 1
	return result        // 1
}

// O(1)
func (Queue *ArrayQueue) Head() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.children[0] // 1
}

// O(1)
func (Queue *ArrayQueue) Tail() interface{} {
	if Queue.IsEmpty() { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.children[Queue.lastPosition-1] // 1
}

//-------------Instanciador del Queue, y enlace de esta con la interfaz IQueue
// O(1)
func NewArrayQueue() *ArrayQueue {
	var ArrayQueue *ArrayQueue = &ArrayQueue{ // 1
		children:     make([]interface{}, 10), // 1
		lastPosition: 0,
	}
	implementIQueue(ArrayQueue) // 1
	return ArrayQueue           // 1
}

//-------------------------------------------------------------------------------------------
