package main

import (
	"fmt"
	// bq "algoritmos/Busquedas"
	// ex "algoritmos/Ejercicios"
	// or "algoritmos/Ordenamiento"
	colas "algoritmos/EstructurasDeDatos/Colas"
)

func main() {
	fmt.Println("Start")

	queue := colas.NewArrayQueue()
	queue.EnQueue(5)
	queue.EnQueue(6)
	fmt.Println("El primer elemento es:", queue.Head())
	fmt.Println("La cola es :", queue.Tail())

}
