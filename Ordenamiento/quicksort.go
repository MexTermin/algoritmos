package ordenamiento

// O(n log n)
func Quicksort(array []int, first int, last int) []int {
	i := first
	j := last
	mid := (array[i] + array[j]) / 2 // 1
	for i < j { // n
		for array[i] < mid { // n (log n)
			i += 1// n (log n)
		}
		for array[j] > mid { // n (log n)
			j -= 1// n (log n)
		}
		if i <= j { // 1
			aux := array[j] // 1
			array[j] = array[i] // 1
			array[i] = aux // 1
			i += 1 // 1
			j -= 1 // 1
		}
	}

	if first < j {
		array = Quicksort(array, first, j) // 1t(n log n)
	}
	if last > i {
		array = Quicksort(array, i, last) // 1t(n log n)
	}
	return array
}

/*
	1t(n log n) + O(1) =  O(n log n)
*/
