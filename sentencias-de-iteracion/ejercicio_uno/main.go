// Realizar un algoritmo para determinar el mayor de 3 números.
// Y luego para determinar el mayor de 10 numeros.
package main

import "fmt"

func main() {
	tresNumeros := []int{3, 23, 55}
	mayorDeTres := obtenerNumeroMayor(tresNumeros)

	if mayorDeTres != nil {
		fmt.Println("El número mayor de 3 numeros es", *mayorDeTres)
	} else {
		fmt.Println("El conjunto de numeros esta vacio")
	}

	diezNumeros := []int{5, 77, 36, 54, 112, 11, 26, 40, 34, 0}
	mayorDeDiez := obtenerNumeroMayor(diezNumeros)

	if mayorDeDiez != nil {
		fmt.Println("El número mayor de 10 numeros es", *mayorDeDiez)
	} else {
		fmt.Println("El conjunto de numeros esta vacio")
	}
}

func obtenerNumeroMayor(nums []int) *int {
	if len(nums) == 0 {
		return nil
	}
	mayor := nums[0]
	for index := 1; index < len(nums); index++ {
		if mayor < nums[index] {
			mayor = nums[index]
		}
	}
	return &mayor
}
