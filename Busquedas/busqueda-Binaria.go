package busquedas

/*

	La Búsqueda Binaria: compara si el valor buscado
	está en la mitad superior o inferior. En la que esté,
	subdivido nuevamente, y así sucesivamente hasta
	encontrar el valor.

*/

// O(log n)
func BusquedaBinaria(array []int, elemento int) int {
	min := 0
	max := len(array) - 1
	for min <= max { // n/2
		mid := (min + max) / 2 // 1
		if array[mid] == elemento {
			return mid // 1
		}
		if array[mid] < elemento {
			min = mid - 1 // 1
		} else if array[mid] > elemento {
			max = mid + 1 //1
		}

	}
	return -1 // 1
}
