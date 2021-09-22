package ejercicios

func CenturyFromYear(year int) int {
	if year < 100 { // 1
		return 1
	}
	if year%100 == 0 {
		return int(year / 100) // 1
	}
	return int(year/100) + 1 // 1
}