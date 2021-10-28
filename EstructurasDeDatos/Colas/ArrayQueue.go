package EstructurasDeDatos

import utils "algoritmos/EstructurasDeDatos/Utils"

type ArrayQueue struct {
	children []interface{}
	len      int
}

// O(1)
func (Q *ArrayQueue) IsEmpty() bool {
	return Q.len == 0 // 1
}

// O(n)
func (Q *ArrayQueue) EnQueue(value interface{}) {
	if Q.len == len(Q.children) { // 1
		Q.children = utils.ResizeArray(Q.children, Q.len*2) // n
	}
	Q.children[Q.len] = value // 1
	Q.len++                   //1
}

// O(n)
func (Q *ArrayQueue) DeQueue() interface{} {
	if Q.len == 0 { // 1
		return nil // 1
	}
	result := Q.children[0]     // 1
	for i := range Q.children { // n
		Q.children[i] = Q.children[i+1] // 1
	}
	if len(Q.children)/2 > Q.len { // 1
		Q.children = utils.ResizeArray(Q.children, Q.len*2) // n
	}
	Q.len--       // 1
	return result // 1
}

// O(1)
func (Q *ArrayQueue) Head() interface{} {
	if Q.len == 0 { // 1
		return nil // 1
	}
	return Q.children[0] // 1
}

// O(1)
func (Q *ArrayQueue) Tail() interface{} {
	if Q.len == 0 { // 1
		return nil // 1
	}
	return Q.children[Q.len-1] // 1
}

//-------------Instanciador del Queue, y enlace de esta con la interfaz IQueue
// O(1)
func NewArrayQueue() *ArrayQueue {
	var ArrayQueue *ArrayQueue = &ArrayQueue{ // 1
		len:      0,
		children: make([]interface{}, 10), // 1
	}
	implementIQueue(ArrayQueue) // 1
	return ArrayQueue           // 1
}

//-------------------------------------------------------------------------------------------
