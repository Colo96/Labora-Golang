// Desarrollar un algoritmo para hallar el mínimo común múltiplo (abreviado como mcm) entre dos números naturales.
// El mínimo común múltiplo entre dos números es el menor número que es divisible por ambos.
package main

import "fmt"

func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mcm(a, b int) int {
	return (a * b) / mcd(a, b)
}

func main() {
	var num1, num2 int

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)

	fmt.Printf("El MCM de %d y %d es: %d\n", num1, num2, mcm(num1, num2))
}
