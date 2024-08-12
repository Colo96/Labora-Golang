// Escribir un algoritmo que determine si un número es par.
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un número")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("El numero", num, "es par")
	} else {
		fmt.Println("El numero", num, "es impar")
	}
}
