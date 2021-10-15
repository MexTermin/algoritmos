package EstructurasDeDatos

// import "fmt"

type IList interface {
	Add(value interface{})
	Delete(index int)
	Insert(index int, value interface{})
	Get(index int) interface{}
	Lenght() int
}
