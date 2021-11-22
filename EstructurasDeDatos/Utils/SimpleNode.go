package Utils

type SimpleNode struct {
	value interface{}
	next  *SimpleNode
}

func (Node *SimpleNode) SetNext(value *SimpleNode) {
	Node.next = value
}
func (Node *SimpleNode) GetNext() *SimpleNode {
	return Node.next
}
func (Node *SimpleNode) SetValue(value interface{}) {
	Node.value = value
}
func (Node *SimpleNode) GetValue() interface{} {
	return Node.value
}
