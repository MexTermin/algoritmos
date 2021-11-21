package EstructurasDeDatos

import utils "algoritmos/EstructurasDeDatos/Utils"

type Lista struct {
	len      int
	children []interface{}
}

// Agrega un elemento al final de la lista
// O(n)
func (Lista *Lista) Add(value interface{}) {
	if len(Lista.children) == Lista.len { // 1
		Lista.children = utils.ResizeArray(Lista.children, Lista.len * 2, Lista.len - 1) // n
	}
	Lista.children[Lista.len] = value // 1
	Lista.len++                   // 1
}

// Elimina el elemento que se encuentra en el index dado por el parametro
// O(n)
func (Lista *Lista) Delete(index int) {
	if index >= Lista.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}
	for i := index; i < Lista.len - 1; i++ { // n
		Lista.children[i] = Lista.children[i + 1] // 1
	}
	if len(Lista.children) / 2 > Lista.len && Lista.len > 10 { // 1
		Lista.children = utils.ResizeArray(Lista.children, Lista.len * 2, Lista.len - 1) // n
	}
	Lista.len-- // 1
}

// Inserta un valor en el index dado por parametro
// O(n)
func (Lista *Lista) Insert(index int, value interface{}) {
	if index >= Lista.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}
	if len(Lista.children) == Lista.len { // 1
		Lista.children = utils.ResizeArray(Lista.children, Lista.len * 2, Lista.len - 1) // n
	}
	for i := Lista.len; i > index; i-- { // n
		Lista.children[i] = Lista.children[i - 1] // 1
	}
	Lista.children[index] = value // 1
	Lista.len++                   // 1
}

// Retorna el valor que se encuentra en la posicion dada
// O(1)
func (Lista *Lista) Get(index int) interface{} {
	if index >= Lista.len || index < 0 { // 1
		panic("Índice fuera de rango") // 1
	}
	return Lista.children[index] // 1
}

// O(1)
func (Lista *Lista) Lenght() int {
	return Lista.len // 1
}

// O(1)
func (Lista *Lista) Value() []interface{} {
	return Lista.children[:Lista.len] // 1
}

//-------------Implementación de la estructura de la lista y enlace de esta con la interfaz
// O(1)
func NewLista() *Lista {
	var ListaStruc *Lista = &Lista{ // 1
		len:      0,                       // 1
		children: make([]interface{}, 10), // 1
	}
	implementInterface(ListaStruc) // 1
	return ListaStruc              // 1
}

//-------------------------------------------------------------------------------------------
