// Desarrollar un algoritmo para hallar el máximo común divisor (abreviado como mcd) entre dos números naturales.
// El máximo común divisor entre dos números es el mayor número que divide a ambos.
package main

import "fmt"

func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)

	fmt.Printf("El MCD de %d y %d es: %d\n", num1, num2, mcd(num1, num2))
}
