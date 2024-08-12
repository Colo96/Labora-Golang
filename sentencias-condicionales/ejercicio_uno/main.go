// Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un número")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("El numero", num, "es mayor a cero")
	} else if num < 0 {
		fmt.Println("El numero", num, "es menor a cero")
	} else {
		fmt.Println("El numero", num, "es cero")
	}
}
