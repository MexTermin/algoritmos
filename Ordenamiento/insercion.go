package ordenamiento

// O(n²)
func InsertionSort(array []int) []int {
	for index := 0; index <= len(array)-1; index++ { // n
		elemento := array[index]            // 1
		j := index - 1                      // 1
		for j >= 0 && array[j] > elemento { // n ( 0 ... + n)
			array[j+1] = array[j] // 1
			j--                   // 1
		}
		array[j+1] = elemento // 1
	}
	return array
}
