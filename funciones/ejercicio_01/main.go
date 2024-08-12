// Realice un algoritmo para determinar si un número es perfecto.
// Un número es perfecto cuando la suma de sus divisores propios es igual al número.
// Los divisores propios de un número son todos sus divisores sin contar el mismo número.
package main

import "fmt"

func main() {
	fmt.Println("Ingrese un numero:")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("El numero ingresado es invalido:", err)
	}

	if esPerfecto(num) {
		fmt.Println(num, "es un número perfecto")
	} else {
		fmt.Println(num, "no es un número perfecto")
	}
}

func esPerfecto(num int) bool {
	suma := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			suma += i
		}
	}
	return suma == num
}
