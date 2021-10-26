package EstructurasDeDatos

// import "fmt"

type IList interface {
	Add(value interface{})
	Delete(index int)
	Insert(index int, value interface{})
	Get(index int) interface{}
	Lenght() int
}

// O(1)
func implementInterface(L IList) IList { // 1
	return L // 1
}
