package ordenamiento

func OrdenamientoInsercion(array []int) []int {
	for index := 0; index <= len(array)-1; index++ { // n
		elemento := array[index]            // 1
		j := index - 1                      // 1
		for j >= 0 && array[j] > elemento { // n (n) / 2
			array[j+1] = array[j] // 1
			j--                   // 1
		}
		array[j+1] = elemento // 1
	}
	return array
}
