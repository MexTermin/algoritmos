package EstructurasDeDatos

// O(1)
func Addicion(lista []int, value int) []int {
	result := append(lista, value) // append a number in the end os slice. O(1)
	return result
}

// O(n)
func Insertion(lista []int, index int, value int) []int {
	lista = append(lista[:index+1], lista[index:]...) // O(n)
	lista[index] = value                              // 1
	return lista                                      // 1
}

//Elimina del slice el elemento del indice colocado, cambiando el elemento dado por el que esta al final
// del arreglo, los lo que siempre lo eliminaraen 2 pasos, pero el orden del arreglo se puede ver afectado
// Time Complexity: O(1)
func RemoveByIndexWithoutOrden(lista []int, index int) []int {
	lista[len(lista)-1], lista[index] = lista[index], lista[len(lista)-1] // O(1)
	return lista[:len(lista)-1]                                           // O(1)
}

// Elimina el elemento del arreglo preservando el orden que se tenia desde el comienzo
// Time Complexity: O(n)
func RemoveByIndex(lista []int, index int) []int {
	return append(lista[:index], lista[index+1:]...) // O(n)
}

// O(n)
func GetIndex(list []int, element int) int {
	for i, v := range list { // O(n)
		if v == element { // O(n)
			return i // O(n)
		}
	}
	return -1 // 1
}


