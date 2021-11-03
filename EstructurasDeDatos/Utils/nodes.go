package utils

type NodeSimple struct {
	Value interface{}
	Next  *NodeSimple
}

type NodeDouble struct {
	Value interface{}
	Prev  *NodeDouble
	Next  *NodeDouble
}
