package ordenamiento

// O(n²)
func Quicksort(array []int, first int, last int) []int {
	i := first
	j := last
	mid := (array[i] + array[j]) / 2
	for i < j { // n
		for array[i] < mid { // n(n)
			i += 1
		}
		for array[j] > mid { // n(n)
			j -= 1
		}
		if i <= j { // 1
			aux := array[j]
			array[j] = array[i]
			array[i] = aux
			i += 1
			j -= 1
		}
	}

	if first < j {
		array = Quicksort(array, first, j) // 1t(log n)
	}
	if last > i {
		array = Quicksort(array, i, last) // 1t(log n)
	}

	return array
}

/*
	1t(log n) * O(n²) =  O(n²)
*/
