package main

import (
	"fmt"
	// bq "algoritmos/Busquedas"
	// ex "algoritmos/Ejercicios"
	// "algoritmos/Ordenamiento"
	colas "algoritmos/EstructurasDeDatos/Colas"
)

func main() {
	fmt.Println("Start")

	queue := colas.NewLinkedQueue()
	queue.EnQueue(1)
	queue.EnQueue(1)
	queue.EnQueue(1)
}
