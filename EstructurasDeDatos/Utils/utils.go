package utils

// O(n)
func ResizeArray(oldArray []interface{}, size int) []interface{} {
	newArray := make([]interface{}, size) // 1
	for i, v := range oldArray {          // n
		newArray[i] = v // 1
	}
	return newArray // 1
}

type NodeSimple struct {
	Value interface{}
	Next  *NodeSimple
}

type NodeDouble struct {
	Value interface{}
	Prev  *NodeDouble
	Next  *NodeDouble
}
