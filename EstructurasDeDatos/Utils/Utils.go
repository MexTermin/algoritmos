package Utils

// Time Complexity: O(n)
func ResizeArray(oldArray []interface{}, size int, endFill int) []interface{} {
	newArray := make([]interface{}, size) // 1
	for index := 0 ; index <= endFill; index++ {          // n
		newArray[index] = oldArray[index] // 1
	}
	return newArray // 1
}
