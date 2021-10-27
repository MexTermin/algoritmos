package EstructurasDeDatos

type IQueue interface {
	isEmpty() bool
	enQueue(interface{})
	deQueue() interface{}
	head() interface{}
	tail() interface{}
}

// O(1)
func ImplementIQueue(L IQueue) IQueue { // 1
	return L // 1
}
