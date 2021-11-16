package EstructurasDeDatos

type IList interface {
	Add(value interface{})
	Delete(index int)
	Insert(index int, value interface{})
	Get(index int) interface{}
	Lenght() int
}

// O(1)
func implementInterface(Lista IList) IList { // 1
	return Lista // 1
}
