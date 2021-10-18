package main

import (
	"fmt"

	// bq "algoritmos/Busquedas"
	// ex "algoritmos/Ejercicios"
	// or "algoritmos/Ordenamiento"
	etd "algoritmos/EstructurasDeDatos"
)

func main() {
	fmt.Println("Start")

	// slice := []int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}
	// leng := len(slice) - 1

	// fmt.Println("Busqueda binaria Recursiva: ", ex.Binary_search_recursiva(slice, 2, 0, 5))
	// fmt.Println("Busqueda Secuencial: ", bq.BusquedaElemento(slice, 2))
	// fmt.Println("Busqueda binaria: ", bq.BusquedaBinaria(slice, 3))

	// fmt.Println("Ordenamiento Burbuja: ", or.OrdenamientoBurbuja(slice))
	// fmt.Println("Ordenamiento Por seleccion: ", or.OrdenamientoSeleccion(slice))
	// fmt.Println("Ordenamiento Por Inserccion: ", or.OrdenamientoInsercion(slice))
	// fmt.Println("Ordenamiento Rapido: ", or.Quicksort(slice, 0, 9))
	// fmt.Println("Ordenamiento por Mezcla: ", or.OrdenamientoMerge(slice, 0, 8))
	// fmt.Println("Ordenamiento por pila: ", or.HeapSort(slice))

	//--------------------------Estructuras de datos------------------------------
	//------Listas
	lista := etd.NewLista()
	lista.Add(5)
	lista.Insert(0, "yael")
	lista.Delete(0)
	fmt.Println("la longitud de la lista es: ", lista.Lenght())
	fmt.Println("La lista es: ", lista)

	// Linked List
	Linked := etd.NewLinkedList()

	Linked.Add(5)
	Linked.Add(9)
	Linked.Add(7)

	Linked.Insert(1, 6)

	Linked.Delete(2)

	fmt.Println(Linked.Get(0))
	fmt.Println(Linked.Lenght())

	Linked.Display()

}
