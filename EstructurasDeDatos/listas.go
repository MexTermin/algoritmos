package EstructurasDeDatos

type Lista struct {
	len      int
	children []interface{}
}

// Se agrega un elemento al final de la lista
// O(n)
func (L *Lista) Add(value interface{}) {
	if len(L.children) == L.len {
		newSlice := make([]interface{}, L.len*2) // 1
		for i, v := range L.children {           // n
			newSlice[i] = v // 1
		}
		newSlice[len(newSlice)-1] = value // 1
		L.children = newSlice             // 1
	}
	L.children[L.len] = value // 1
	L.len++                   // 1
}

// Se elimina el elemento que se encuentra en el index pasado por parametro
// O(n)
func (L *Lista) Delete(index int) {
	if index >= L.len || index < 0 {
		panic("Indice fuera de rango")
	}
	newSlice := make([]interface{}, len(L.children)) // 1
	var deleted int = 0                              // 1
	for i := L.len; i > 0; i-- {                     // n
		if i == index {
			deleted = 1 // 1
		}
		newSlice[i-1] = L.children[i-deleted] //1
	}
	L.len--               // 1
	L.children = newSlice // 1
}

// Se inserta un valor en el index dado por parametro
// O(n)
func (L *Lista) Insert(index int, value interface{}) {
	if index >= L.len || index < 0 {
		panic("Indice fuera de rango")
	}
	var newSlice []interface{} = L.children // 1

	if len(L.children) == L.len {
		newSlice = make([]interface{}, len(L.children)*2) // 1
	}
	for i := L.len - 1; i >= index; i-- { // n
		if i == index { // 1
			newSlice[i+1] = newSlice[i] // 1
			newSlice[i] = value         // 1
			break                       // 1
		}
		newSlice[i+1] = L.children[i] //1
	}
	L.len++               // 1
	L.children = newSlice // 1
}

// Retorna el valor que se encuentra en la posicion dada
// O(1)
func (L *Lista) Get(index int) interface{} {
	if index >= L.len || index < 0 {
		panic("Indice fuera de rango")
	}
	return L.children[index] // 1
}

// O(1)
func (L *Lista) Lenght() int {
	return L.len // 1
}

// O(1)
func (L *Lista) Value() []interface{} {
	return L.children[:L.len] // 1
}

//-------------Implementacion de la estructura de la lista, y enlace de esta con la interfaz

func NewLista() *Lista {
	var ListaStruc *Lista = &Lista{
		len:      0,
		children: make([]interface{}, 10),
	}
	implementInterface(ListaStruc)
	return ListaStruc
}

//-------------------------------------------------------------------------------------------
