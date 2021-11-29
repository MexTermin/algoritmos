package Colas

import (
	utils "algoritmos/EstructurasDeDatos/Utils"
	"fmt"
)

type PriorityLinkedQueue struct {
	LinkedQueue
}

func (Queue *PriorityLinkedQueue) EnQueue(args interface{}) {
	var (
		array    []interface{} = args.([]interface{}) // 1
		value    interface{}   = array[0]             // 1
		priority int           = array[1].(int)       // 1
	)

	node := &utils.SimpleNode{} // 1
	node.SetValue(value)        // 1
	node.SetPriority(priority)  // 1

	if Queue.IsEmpty(){
		Queue.head = node // 1
		Queue.tail = node // 1
	} else {
		nodeActual := Queue.head          // 1
		for nodeActual.GetNext() != nil { // 1
			if nodeActual.GetPriority() > priority { // 1
				nodeActual = node.GetNext() // 1
				continue // 1
			}
			break // 1
		}

		node.SetNext(nodeActual.GetNext()) // 1
		nodeActual.SetNext(node)           // 1
	}
}

func (Queue *PriorityLinkedQueue) Print() {
	nodo := Queue.head
	for nodo != nil {
		fmt.Print(" -> ", nodo.GetValue())
		nodo = nodo.GetNext()
	}
	fmt.Println("")
}

func NewPriorityLinkedQueue() *PriorityLinkedQueue {
	queue := &PriorityLinkedQueue{}
	implementIQueue(queue)
	return queue
}
