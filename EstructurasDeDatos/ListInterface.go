package EstructurasDeDatos

import "fmt"

type IList interface {
	Add(value interface{})
	Delete(value interface{})
	Insert(value interface{}) string
	Get(value interface{}) string
	Lenght() int
}

type Lista struct {
	len      int
	children []interface{}
}

func (L *Lista) Add(value interface{}) {
	newSlice := make([]interface{}, len(L.children)+1)
	for i, v := range L.children {
		newSlice[i] = v
	}
	newSlice[len(newSlice)-1] = value
	L.len++
	L.children = newSlice
}

func (L *Lista) Delete(value interface{}) {
	newSlice := make([]interface{}, len(L.children)-1)
	var deleted int = 0
	for i, v := range L.children {
		if v != value {
			fmt.Println("el ", v, "es", value, "del index", i)
			newSlice[i-deleted] = v
			continue
		}
		deleted = 1
	}
	L.len--
	L.children = newSlice
}

func (L *Lista) Insert(value interface{}) string {
	return "insert"
}

func (L *Lista) Get(value interface{}) string {
	return "get"
}

func (L *Lista) Lenght() int {
	return L.len
}

func implementInterface(L IList) IList {
	return L
}

func NewLista() *Lista {
	var ListaStruc *Lista = new(Lista)
	ListaStruc.len = 0
	ListaStruc.children = make([]interface{}, 0)
	implementInterface(ListaStruc)
	return ListaStruc
}
