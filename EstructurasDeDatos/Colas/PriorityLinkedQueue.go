package Colas

import utils "algoritmos/EstructurasDeDatos/Utils"

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

	if Queue.IsEmpty() { // 1
		Queue.head = node // 1
		Queue.tail = node // 1
	} else { // 1
		var nodeActual *utils.SimpleNode         // 1
		if Queue.tail.GetPriority() < priority { // 1
			Queue.tail.SetNext(node) // 1
			Queue.tail = node        // 1
		} else {
			nodeActual = Queue.head                   // 1
			for nodeActual.GetPriority() > priority { // 1
				if nodeActual.GetNext() != nil { // 1
					nodeActual = node.GetNext() // 1
				} else { // 1
					break // 1
				}
			}
			node.SetNext(nodeActual.GetNext()) // 1
			nodeActual.SetNext(node)           // 1
		}
	}
}

//-------------Implementación de la estructura de colas con prioridad y enlace de esta con la interfaz Queue
// O(1)
func NewPriorityLinkedQueue() *PriorityLinkedQueue {
	queue := &PriorityLinkedQueue{} // 1
	implementIQueue(queue) // 1
	return queue // 1
}
