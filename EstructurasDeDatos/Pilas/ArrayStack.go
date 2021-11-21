package pilas

import utils "algoritmos/EstructurasDeDatos/utils"

type Stack struct {
	children []interface{} // 1
	size     int // 1
}

// O(1)
func (Stack *Stack) IsEmpty() bool {
	return Stack.size == 0 // 1
}

// O(n)
func (Stack *Stack) Push(value interface{}) {
	if Stack.size == len(Stack.children) { // 1
		Stack.children = utils.ResizeArray(Stack.children, len(Stack.children)*2) // n
	}

	Stack.children[Stack.size] = value // 1
	Stack.size++ // 1
}

// O(n)
func (Stack *Stack) Pop() interface{}{
	if Stack.IsEmpty(){ // 1
		panic("La pila está vacía") // 1
	}

	if  len(Stack.children) / 2 > Stack.size && Stack.size > 10 { // 1
		Stack.children = utils.ResizeArray(Stack.children, len(Stack.children)*2) // n
	}

	var result interface{} = Stack.children[Stack.size - 1] // 1
	Stack.children[Stack.size - 1] = nil // 1
	return result // 1
}

// O(1)
func (Stack *Stack) Top() interface{}{
	if Stack.IsEmpty(){ // 1
		panic("la pila está vacía") // 1
	}
	return Stack.children[Stack.size - 1] // 1
}

// O(1)
func NewStack() *Stack {
	var stack *Stack = &Stack{ // 1
		children: make([]interface{}, 10), // 1
		size:     0,                       // 1
	} // 1

	ImplementStack(stack) // O(1)
	return stack          // 1
}
