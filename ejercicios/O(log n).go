package ejercicios

func Binary_search(array []int, elemento int) int {
	min := 0
	max := len(array) - 1 // 1
	index := 0

	for min <= max { // n / 2
		min = (max + min) / 2
		if array[index] < elemento {
			min = index + 1 // 1
		} else if array[index] > elemento {
			max = index - 1 // 1
		} else {
			return index
		}
	}
	return -1 // 1
}

func Binary_search_recursiva(arreglo []int, busqueda int, primerElemento int, UltimoElemento int) int {
	if primerElemento > UltimoElemento { // 1
		return -1
	}
	mid := (primerElemento + UltimoElemento) / 2 // 1
	elementMid := arreglo[mid] // 1
	if elementMid == busqueda {
		return mid // 1
	}
	if busqueda < elementMid {
		return Binary_search_recursiva(arreglo, busqueda, primerElemento, mid-1) // T(n) / 2
	} else {
		return Binary_search_recursiva(arreglo, busqueda, mid+1, UltimoElemento) // T(n) / 2
	}
}
