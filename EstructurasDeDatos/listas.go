package EstructurasDeDatos

type Lista struct {
	len      int
	children []interface{}
}

// Se agrega un elemento al final de la lista
// O(n)
func (L *Lista) Add(value interface{}) {
	newSlice := make([]interface{}, len(L.children)+1) // 1
	for i, v := range L.children { // n
		newSlice[i] = v // n
	}
	newSlice[len(newSlice)-1] = value // 1
	L.len++ // 1
	L.children = newSlice // 1
}

// Se elimina el elemento que se encuentra en el index pasado por parametro
// O(n)
func (L *Lista) Delete(index int) {
	newSlice := make([]interface{}, len(L.children)-1) // 1
	var deleted int = 0 // 1
	for i, v := range L.children { // n
		if i != index { // n
			newSlice[i-deleted] = v // n
			continue
		}
		deleted = 1 // n
	}
	L.len-- // 1
	L.children = newSlice // 1
}

// Se inserta un valor en el index dado por parametro
// O(n)
func (L *Lista) Insert(index int, value interface{}) {
	newSlice := make([]interface{}, len(L.children)+1) // 1
	var incrementIndex = 0
	for i, v := range L.children { // n
		if i == index { // n
			newSlice[i] = value // n
			newSlice[i+1] = v // n
			incrementIndex = 1 // n
			continue // n
		}
		newSlice[i+incrementIndex] = v // n
	}
	L.len++ // 1
	L.children = newSlice // 1
}

// Retorna el valor que se encuentra en la posicion dada
// O(n)
func (L *Lista) Get(index int) interface{} {
	for i, v := range L.children { // n
		if i == index { // n
			return v // n
		}
	}
	return -1 // 1
}

// O(1)
func (L *Lista) Lenght() int {
	return L.len // 1
}

//-------------Implementacion de la estructura de la lista, y enlace de esta con la interfaz
func implementInterface(L IList) IList {
	return L
}

func NewLista() IList {
	var ListaStruc *Lista = &Lista{
		len:      0,
		children: make([]interface{}, 0),
	}
	return implementInterface(ListaStruc)
}
//-------------------------------------------------------------------------------------------