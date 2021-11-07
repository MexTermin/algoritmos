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
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	// fmt.Println("La cabeza es: ", queue.Head())
	// fmt.Println("La cabeza es: ", queue.Tail())
	fmt.Println("is Empty: ", queue.IsEmpty())
}
