package main

import (
	"fmt"
	// bq "algoritmos/Busquedas"
	// ex "algoritmos/Ejercicios"
	// "algoritmos/Ordenamiento"
	"algoritmos/EstructurasDeDatos/Colas"
	// "algoritmos/EstructurasDeDatos/Listas"
	// pilas "algoritmos/EstructurasDeDatos/Pilas"
)

func main() {
	fmt.Println("Start")
	cola := Colas.NewPriorityLinkedQueue()
	cola.EnQueue([]interface{}{"1", 1})
	cola.EnQueue([]interface{}{"2", 2})
	cola.EnQueue([]interface{}{"3", 2})
	cola.EnQueue([]interface{}{"4", 2})
	cola.Print()
}
