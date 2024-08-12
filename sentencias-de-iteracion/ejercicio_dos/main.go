// Realice un algoritmo para sumar los primeros 3 números naturales.
// Y luego un algoritmo para sumar los primeros 10 números naturales.
package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sumaTres, err := sumarNumeros(numeros, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La suma de los primeros tres números naturales es", sumaTres)
	}

	sumaDiez, err := sumarNumeros(numeros, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La suma de los 10 números naturales es", sumaDiez)
	}
}

func sumarNumeros(numeros []int, cantidad int) (int, error) {
	if cantidad > len(numeros) {
		return 0, fmt.Errorf("no hay suficientes números en el conjunto para sumar los primeros %d", cantidad)
	}

	suma := 0
	for _, numero := range numeros[:cantidad] {
		suma += numero
	}
	return suma, nil
}
