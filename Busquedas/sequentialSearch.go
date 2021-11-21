package busquedas

/*

	Buqueda Secuencial: Consiste en ir comparando el elemento que se busca con cada elemento del arreglo hasta cuando se encuentra.

*/

// retorna el index del elemento que deseas buscar, si no lo encuentra regresa -1
// O(n)
func SearchElement(array []int, number int) int {

	for index, value := range array { // n
		if value == number { // 1
			return index
		}
	}
	return -1
}

// retorna el elemento mas pequeño del arreglo
// O(n)
func SmallestSearch(array []int, number int) int {
	menor := 0
	for _, value := range array { // n
		if value <= menor { // 1
			menor = value // 1
		}
	}
	return menor // 1
}

// retorna el elemento mas grande del arreglo
// O(n)
func HigherSearch(array []int, number int) int {
	mayor := 0
	for _, value := range array { // n
		if value >= mayor { // 1
			mayor = value // 1
		}
	}
	return mayor // 1
}
