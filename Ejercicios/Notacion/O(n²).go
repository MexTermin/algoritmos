package Ejercicios

func QuickSort(array []int, first int, last int) []int {
	i := first // 1
	j := last // 1
	mid := (array[i] + array[j]) / 2
	for i < j { // n
		for array[i] < mid { // n(n)
			i += 1 // 1
		}
		for array[j] > mid { // n(n)
			j -= 1 // 1
		}
		if i <= j { // 1
			aux := array[j] // 1
			array[j] = array[i] // 1
			array[i] = aux // 1
			i += 1 // 1
			j -= 1 // 1
		}
	}

	if first < j { // 1
		array = QuickSort(array, first, j) // 1t(n / 2)
	}
	if last > i { // 1
		array = QuickSort(array, i, last) // 1t(n / 2)
	}

	return array // 1
}
