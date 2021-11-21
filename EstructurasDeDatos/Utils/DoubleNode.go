package utils

type DoubleNode struct {
	value interface{}
	prev  *DoubleNode
	next  *DoubleNode
}

func (Node *DoubleNode) SetPrev(value *DoubleNode) {
	Node.prev = value
}
func (Node *DoubleNode) GetPrev() *DoubleNode {
	return Node.prev
}
func (Node *DoubleNode) SetNext(value *DoubleNode) {
	Node.next = value
}
func (Node *DoubleNode) GetNext() *DoubleNode {
	return Node.next
}
func (Node *DoubleNode) SetValue(value interface{}) {
	Node.value = value
}
func (Node *DoubleNode) GetValue() interface{} {
	return Node.value
}
