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

	slice := []int{1, 8, 4, 3, 26, 45, 78, 10, 23, 98}
	leng := len(slice) - 1

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
	//--Listas
	fmt.Println("Insertion: ", etd.Insertion(slice, 9, 15))
	fmt.Println("Remove an element by index: ", etd.RemoveByIndexWithoutOrden(slice, leng))
	fmt.Println("Remove an element by index: ", etd.RemoveByIndex(slice, leng))
	fmt.Println("Obtener Longitud: ", len(slice)) // O(1)

	//--Linked List
	Linked := &etd.Linked{} // make a new intance

	// añado elementos
	Linked.Push(5)
	Linked.Push(3)
	Linked.Push(2)

	// muestro los nodos enlazados
	Linked.Display()

	// imprimo lalongitud de la lista
	fmt.Println(Linked.Len())

	// busco el nodo con el valor 3
	node := Linked.FindNode(3)
	Linked.InsertAfter(node, 9)

	node = Linked.FindNode(9)
	Linked.InsertBefore(node, 1)

	Linked.Display()

}
