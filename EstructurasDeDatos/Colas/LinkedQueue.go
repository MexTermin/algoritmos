package EstructurasDeDatos

type Node struct {
	value interface{}
	next  *Node
}

type LinkedQueue struct {
	head *Node
	tail *Node
}

// O(1)
func (Queue *LinkedQueue) IsEmpty() bool {
	return Queue.head != nil // 1
}

// O(1)
func (Queue *LinkedQueue) EnQueue(value interface{}) {
	var node = &Node{ // 1
		value: value, // 1
		next:  nil,   // 1
	}
	if Queue.head == nil || Queue.tail == nil { // 1
		Queue.head = node // 1
		Queue.tail = node // 1
	} else {
		Queue.tail.next = node // 1
		Queue.tail = node      // 1
	}
}

// O(1)
func (Queue *LinkedQueue) DeQueue() interface{} {
	if Queue.head == nil { // 1
		panic("La cola esta vacia") // 1
	}
	var result interface{} = Queue.head.value // 1
	if Queue.head.next != nil {               // 1
		Queue.head = Queue.head.next // 1
	} else {
		Queue.head = nil // 1
		Queue.tail = nil // 1
	}
	return result // 1
}

// O(1)
func (Queue *LinkedQueue) Head() interface{} {
	if Queue.head == nil { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.head.value //1
}

// O(1)
func (Queue *LinkedQueue) Tail() interface{} {
	if Queue.tail == nil { // 1
		panic("La cola esta vacia") // 1
	}
	return Queue.tail.value //1
}

//-------------Implementacion de la estructura de colas, y enlace de esta con la interfaz Queue
// O(1)
func NewLinkedQueue() *LinkedQueue {
	var ListaStruc *LinkedQueue = &LinkedQueue{ // 1
		head: nil, // 1
		tail: nil, // 1
	}
	implementIQueue(ListaStruc) // 1
	return ListaStruc           // 1
}

//-------------------------------------------------------------------------------------------
