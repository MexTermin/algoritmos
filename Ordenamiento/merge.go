package ordenamiento

// O(n * log(n))
func merge(array []int, left int, mid int, right int) []int {

	// merge(n) = O(12n/2 + 5log(n/2) + o(1)) =  O(n + log(n) + O(1)) = O(n log(n))

	n1 := mid - left + 1 // 1
	n2 := right - mid    // 1

	var i, j int
	var R []int = make([]int, n1) // 1
	var L []int = make([]int, n1) // 1

	for i := 0; i < n1; i++ { // n / 2
		L[i] = array[left+i] // n / 2
	}
	for j := 0; j < n2; j++ { // n / 2
		R[j] = array[mid+1+j] // n / 2
	}

	i = 0 // 1
	j = 0 // 1

	indexToMerge := left   // 1
	for i < n1 && j < n2 { // log(n/2)
		if L[i] <= R[j] { // log(n/2)
			array[indexToMerge] = L[i] // log(n/2)
			i++                        // log(n/2)
		} else {
			array[indexToMerge] = R[j]
			j++
		}
		indexToMerge++ // log(n/2)
	}

	for i < n1 { // n / 2
		array[indexToMerge] = L[i] // n / 2
		i++                        // n / 2
		indexToMerge++             // n / 2
	}

	for j < n2 { // n / 2
		array[indexToMerge] = R[j] // n / 2
		j++                        // n / 2
		indexToMerge++             // n / 2
	}
	return array // 1

}

// O(n log(n))
func MegeSort(array []int, left int, right int) []int {
	if left < right {
		mid := left + (right-left)/2 // 1

		MegeSort(array, left, mid)    // T(n/2)
		MegeSort(array, mid+1, right) // T(n/2)

		merge(array, left, mid, right) // O(n log(n))
	}
	return array
	// recurrencia
	// 2T(n/2) + O(n log(n)) + 1 =  O(n log(n))
}
