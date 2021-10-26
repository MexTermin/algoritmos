package EstructurasDeDatos

type Lista struct {
	len      int
	children []interface{}
}

// O(n)
func resizeArray(oldArray []interface{}, size int) []interface{} {
	newArray := make([]interface{}, size) // 1
	for i, v := range oldArray {          // n
		newArray[i] = v // 1
	}
	return newArray
}

// Se agrega un elemento al final de la lista
// O(n)
func (L *Lista) Add(value interface{}) {
	if len(L.children) == L.len {
		L.children = resizeArray(L.children, L.len*2) // n
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
	for i := index; i < L.len-1; i++ { // n
		L.children[i] = L.children[i+1] //1
	}
	if len(L.children)/2 > L.len && L.len > 10 { // 1
		L.children = resizeArray(L.children, L.len*2) // n
	}
	L.len-- // 1
}

// Se inserta un valor en el index dado por parametro
// O(n)
func (L *Lista) Insert(index int, value interface{}) {
	if index >= L.len || index < 0 {
		panic("Indice fuera de rango")
	}
	if len(L.children) == L.len {
		L.children = resizeArray(L.children, L.len*2) // n
	}
	for i := L.len; i > index; i-- { // n
		L.children[i] = L.children[i-1] //1
	}
	L.children[index] = value
	L.len++ // 1
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
