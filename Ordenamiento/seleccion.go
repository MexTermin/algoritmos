package ordenamiento

// O(n²)
func SelectionSort(array []int) []int {

	for index := 0; index < len(array)-1; index++ { // n
		min := index
		for element := index + 1; element < len(array); element++ { // n (n)
			if array[element] < array[min] { // n (n)
				min = element // n (n)
			}
		}
		if index != min { // 1
			aux := array[index] // 1
			array[index] = array[min] // 1
			array[min] = aux // 1
		}

	}
	return array
}
