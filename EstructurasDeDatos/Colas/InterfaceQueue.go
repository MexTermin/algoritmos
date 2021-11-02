package EstructurasDeDatos

type IQueue interface {
	IsEmpty() bool
	EnQueue(interface{})
	DeQueue() interface{}
	Head() interface{}
	Tail() interface{}
}

// O(1)
func implementIQueue(Queue IQueue) IQueue { // 1
	return Queue // 1
}
