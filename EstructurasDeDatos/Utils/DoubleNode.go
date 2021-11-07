package utils

type NodeDouble struct {
	value interface{}
	prev  *NodeDouble
	next  *NodeDouble
}

func (Node *NodeDouble) SetPrev(value *NodeDouble) {
	Node.prev = value
}
func (Node *NodeDouble) GetPrev() *NodeDouble {
	return Node.prev
}
func (Node *NodeDouble) SetNext(value *NodeDouble) {
	Node.next = value
}
func (Node *NodeDouble) GetNext() *NodeDouble {
	return Node.next
}
func (Node *NodeDouble) SetValue(value interface{}) {
	Node.value = value
}
func (Node *NodeDouble) GetValue() interface{} {
	return Node.value
}
