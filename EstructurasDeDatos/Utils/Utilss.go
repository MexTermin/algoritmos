package utils

// Time Complexity: O(n)
func ResizeArray(oldArray []interface{}, size int) []interface{} {
	newArray := make([]interface{}, size) // 1
	for i, v := range oldArray {          // n
		newArray[i] = v // 1
	}
	return newArray // 1
}
