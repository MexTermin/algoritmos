package Pilas

import utils "algoritmos/EstructurasDeDatos/Utils"

type LinkedStack struct {
	head *utils.SimpleNode // 1
}

// O(1)
func (Stack *LinkedStack) Push(value interface{}) {
	node := &utils.SimpleNode{} // 1
	node.SetValue(value)        // 1
	if !Stack.IsEmpty() {       // 1
		node.SetNext(Stack.head) // 1
	}
	Stack.head = node // 1
}

// O(1)
func (Stack *LinkedStack) Pop() interface{} {
	if Stack.IsEmpty() { // 1
		panic("La pila está vacía") // 1
	}
	var value interface{} = Stack.head.GetValue() // 1
	Stack.head = Stack.head.GetNext()             // 1
	return value                                  // 1
}

// O(1)
func (Stack *LinkedStack) Top() interface{} {
	if Stack.IsEmpty() { // 1
		panic("La pila está vacía") // 1
	}
	return Stack.head.GetValue() // 1
}

// O(1)
func (Stack *LinkedStack) IsEmpty() bool {
	return Stack.head == nil // 1
}

// O(1)
func NewLinkedStack() *LinkedStack {
	var stack *LinkedStack = &LinkedStack{ // 1
		head: nil, // 1
	}
	ImplementStack(stack) // O(1)
	return stack          // 1
}
