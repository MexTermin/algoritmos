package main

import (
	"fmt"
	// bq "algoritmos/Busquedas"
	// ex "algoritmos/Ejercicios"
	// "algoritmos/Ordenamiento"
	// "algoritmos/EstructurasDeDatos/Colas"
	// "algoritmos/EstructurasDeDatos/Listas"
	pilas "algoritmos/EstructurasDeDatos/Pilas"
)

func main() {
	fmt.Println("Start")
	pila := pilas.NewStack()
	pila.Push(1)
	pila.Push(2)
	fmt.Println("Se hizo pop del: ", pila.Pop())
	fmt.Println("El elemento en la sima es: ", pila.Top())
	fmt.Println("Se hizo pop del: ", pila.Pop())

	fmt.Println(pila)
}
