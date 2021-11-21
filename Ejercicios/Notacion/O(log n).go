package Ejercicios

func BinarySearch(array []int, element int) int {
	min := 0 // 1
	max := len(array) - 1 // 1
	index := 0 // 1

	for min <= max { // n / 2
		min = (max + min) / 2
		if array[index] < element { // 1
			min = index + 1 // 1
		} else if array[index] > element { // 1
			max = index - 1 // 1
		} else {
			return index // 1
		}
	}
	return -1 // 1
}

func RecursiveBinarySearch(array []int, search int, firstElement int, lastElement int) int {
	if firstElement > lastElement { // 1
		return -1
	}
	mid := (firstElement + lastElement) / 2 // 1
	elementMid := array[mid] // 1
	if elementMid == search {
		return mid // 1
	}
	if search < elementMid { // 1
		return RecursiveBinarySearch(array, search, firstElement, mid-1) // 1T(n) / 2
	} else {
		return RecursiveBinarySearch(array, search, mid+1, lastElement) // 1T(n) / 2
	}
}
