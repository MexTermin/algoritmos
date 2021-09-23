package ejercicios

import (
    "sort"
)


func Palindrome(cadena string) bool {
	var inverted string = "" // 1
	for i := len(cadena); i > 0; i-- { // n
		inverted += string(cadena[i-1]) // 1
	}
	return inverted == cadena // 1
}

/*
	Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.
	For inputArray = [3, 6, -2, -5, 7, 3], the output should be
	adjacentElementsProduct(inputArray) = 21.

	7 and 3 produce the largest product.
*/
func AdjacentElementsProduct(inputArray []int) int {
	var product int = inputArray[0] * inputArray[1]
	for index := range inputArray[1:] {
		if index == len(inputArray)-1 {
			return product
		}
		multiplicacion := inputArray[index] * inputArray[index+1]
		if multiplicacion >= product {
			product = multiplicacion
		}
	}
	return product
}

func fila(n int) int {
	return 1 + (n-1)*2 // 1
}
/*
Below we will define an n-interesting polygon. Your task is to find the area of a polygon for a given n.

A 1-interesting polygon is just a square with a side of length 1.
An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim,
side by side. You can see the 1-, 2-, 3- and 4-interesting polygons in the picture below.
*/
func ShapeArea(n int) int {
	var result int // 1
	for campo := n - 1; campo > 0; campo-- { // n
		result += fila(campo)
	}
	result = (result * 2) + fila(n) // 1
	return result // 1
}

func MakeArrayConsecutive2(statues []int)int {
    sort.Ints(statues) // O(n)
    var result int // 1
    for i:= range statues{ // O(n)
        if i+1  >= len(statues){ break } // 1
        result += (statues[i+1] - statues[i])-1 // 1
    }
    return result
}