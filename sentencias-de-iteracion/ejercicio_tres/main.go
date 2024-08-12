// Realice un algoritmo para sumar los primeros "n" números naturales, siendo "n" un valor que ingresa el usuario por teclado durante la ejecución del algoritmo
package main

import (
	"fmt"
	"log"
)

func main() {
	var numero int
	fmt.Print("Ingrese un número: ")
	_, err := fmt.Scan(&numero)
	if err != nil || numero < 1 {
		log.Fatalf("Entrada inválida: %v. Por favor, ingrese un número natural positivo.", err)
	}
	fmt.Printf("El usuario ingresó %d y la suma es %d\n", numero, sumarNumeros(numero))
}

func sumarNumeros(numero int) int {
	return numero * (numero + 1) / 2
}
