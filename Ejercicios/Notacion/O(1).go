package Ejercicios

func CenturyFromYear(year int) int {
	if year < 100 { // 1
		return 1 // 1
	}
	if year % 100 == 0 { // 1
		return int(year / 100) // 1
	}
	return int(year / 100) + 1 // 1
}