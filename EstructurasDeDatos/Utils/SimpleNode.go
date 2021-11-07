package utils

type NodeSimple struct {
	value interface{}
	next  *NodeSimple
}

func (Node *NodeSimple) SetNext(value *NodeSimple) {
	Node.next = value
}
func (Node *NodeSimple) GetNext() *NodeSimple {
	return Node.next
}
func (Node *NodeSimple) SetValue(value interface{}) {
	Node.value = value
}
func (Node *NodeSimple) GetValue() interface{} {
	return Node.value
}
