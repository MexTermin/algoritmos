package main

import (
	bq "algoritmos/Busquedas"
	ex "algoritmos/Ejercicios"
	or "algoritmos/Ordenamiento"
	"fmt"
)

func main() {
	fmt.Println("Busqueda binaria Recursiva: ", ex.Binary_search_recursiva([]int{0, 5, 1, 6, 2, 3}, 2, 0, 5))
	fmt.Println("Busqueda Secuencial: ", bq.BusquedaElemento([]int{0, 5, 1, 6, 2, 3}, 2))
	fmt.Println("Busqueda binaria: ", bq.BusquedaBinaria([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	fmt.Println("Ordenamiento Burbuja: ", or.OrdenamientoBurbuja([]int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}))
	fmt.Println("Ordenamiento Por seleccion: ", or.OrdenamientoSeleccion([]int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}))
	fmt.Println("Ordenamiento Por Inserccion: ", or.OrdenamientoInsercion([]int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}))
	fmt.Println("Ordenamiento Rapido: ", or.Quicksort([]int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}, 0, 9))
	fmt.Println("Ordenamiento por Mezcla: ", or.OrdenamientoMerge([]int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}, 0, 8))
}
