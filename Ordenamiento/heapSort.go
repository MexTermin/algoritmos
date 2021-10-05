package ordenamiento

// log n
func Heapify(arr []int, arrayLen int, index int) {
	var largest int = index     // 1
	var left int = 2*index + 1  // 1
	var right int = 2*index + 2 // 1

	if left < arrayLen && arr[left] > arr[largest] { // 1
		largest = left // 1
	}
	if right < arrayLen && arr[right] > arr[largest] { // 1
		largest = right // 1
	}
	if largest != index { // 1
		var temp int = arr[index]       // 1
		arr[index] = arr[largest]       // 1
		arr[largest] = temp             // 1
		Heapify(arr, arrayLen, largest) // T(n * 2 ) + 3(1)) = T(log n) + O(1)
	}
}

//  n * log n
func HeapSort(arr []int) []int {
	var n int = len(arr)                        // 1
	for index := n/2 - 1; index >= 0; index-- { // n / 2
		Heapify(arr, n, index) // n / 2 * log n
	}
	for index := n - 1; index > 0; index-- { // n
		var temp int = arr[0]  // n
		arr[0] = arr[index]    // n
		arr[index] = temp      // n
		Heapify(arr, index, 0) // n * log n
	}
	return arr
} // 2( n / 2 ) * log n + 4(1) = n * log n  * O(1)
