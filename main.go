package main

import (
	ex "algoritmos/Ejercicios"
	bq "algoritmos/Busquedas"
	"fmt"
)

func main() {
	fmt.Println(ex.Binary_search_recursiva([]int{0, 5, 1, 6, 2, 3}, 2, 0, 5))
	fmt.Println(bq.BusquedaElemento([]int{0, 5, 1, 6, 2, 3}, 2))
}
