package EstructurasDeDatos

type Lista struct {
	len      int
	children []interface{}
}

// Se agrega un elemento al final de la lista
// O(n)
func (L *Lista) Add(value interface{}) {
	if len(L.children) < cap(L.children) {
		// L.children[L.len-1] = value
		L.children = append(L.children, value)
	} else {
		newSlice := make([]interface{}, len(L.children)+1, len(L.children)*2) // 1
		for i, v := range L.children {                                        // n
			newSlice[i] = v // 1
		}
		newSlice[len(newSlice)-1] = value // 1
		L.children = newSlice             // 1
	}
	L.len++ // 1
}

// Se elimina el elemento que se encuentra en el index pasado por parametro
// O(n)
func (L *Lista) Delete(index int) {
	if index > L.len {
		panic("Indice fuera de rango")
	} else {
		newSlice := make([]interface{}, len(L.children)-1) // 1
		var deleted int = 0                                // 1
		for i, v := range L.children {                     // n
			if i != index { // 1
				newSlice[i-deleted] = v // 1
				continue
			}
			deleted = 1 // 1
		}
		L.len--               // 1
		L.children = newSlice // 1
	}
}

// Se inserta un valor en el index dado por parametro
// O(n)
func (L *Lista) Insert(index int, value interface{}) {
	if index > L.len {
		panic("Indice fuera de rango")
	} else {

		newSlice := make([]interface{}, len(L.children)+1) // 1
		var incrementIndex = 0
		for i, v := range L.children { // n
			if i == index { // 1
				newSlice[i] = value // 1
				newSlice[i+1] = v   // 1
				incrementIndex = 1  // 1
				continue            // 1
			}
			newSlice[i+incrementIndex] = v // 1
		}
		L.len++               // 1
		L.children = newSlice // 1
	}
}

// Retorna el valor que se encuentra en la posicion dada
// O(1)
func (L *Lista) Get(index int) interface{} {
	if index > L.len {
		panic("Indice fuera de rango")
	} else {
		return L.children[index]
	}
}

// O(1)
func (L *Lista) Lenght() int {
	return L.len // 1
}

//-------------Implementacion de la estructura de la lista, y enlace de esta con la interfaz

func NewLista() *Lista {
	var ListaStruc *Lista = &Lista{
		len:      0,
		children: make([]interface{}, 0, 10),
	}
	implementInterface(ListaStruc)
	return ListaStruc
}

//-------------------------------------------------------------------------------------------
