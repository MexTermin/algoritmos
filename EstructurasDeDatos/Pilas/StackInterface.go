package Pilas

// O(1)
type IStack interface {
	Push(value interface{}) // 1
	Pop() interface{}// 1
	Top() interface{} // 1
	IsEmpty() bool // 1
}

// O(1)
func ImplementStack(stack IStack) IStack{
	return stack // 1
}