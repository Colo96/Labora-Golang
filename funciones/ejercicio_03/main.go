// Escriba un algoritmo para determinar si un número es primo.
// Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo.
package main

import (
	"fmt"
	"math"
)

func pedirNumero(mensaje string) (int, error) {
	var num int
	fmt.Print(mensaje)
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("el número ingresado es inválido: %v", err)
	}
	if num <= 0 {
		return 0, fmt.Errorf("el número debe ser positivo")
	}
	return num, nil
}

func esPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num, err := pedirNumero("Ingrese un número: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if esPrimo(num) {
		fmt.Println(num, "es primo")
	} else {
		fmt.Println(num, "no es primo")
	}
}
