package ordenamiento

// O(n²)
func OrdenamientoBurbuja(array []int) []int {

	for limite := len(array) - 1; limite > 0; limite-- { // n
		for index := 0; index < limite; index++ { // 1 + 2 + 3... + n =  n (n + 1)
			if array[index] > array[index+1] {
				temp := array[index] // 1
				array[index] = array[index+1] // 1
				array[index+1] = temp // 1
			}
		}
	}
	return array // 1
}
