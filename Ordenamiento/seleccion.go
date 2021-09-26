package ordenamiento

func OrdenamientoSeleccion(array []int) []int {

	for index := 0; index < len(array)-1; index++ { // n
		min := index
		for element := index + 1; element < len(array); element++ { // n (n) / 2
			if array[element] < array[min] {
				min = element
			}
		}
		if index != min {
			aux := array[index]
			array[index] = array[min]
			array[min] = aux
		}

	}
	return array
}
