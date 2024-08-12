// Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.
package main

import "fmt"

func main() {
	temp1 := 40
	temp2 := 54
	temp3 := -5

	suma := temp1 + temp2 + temp3
	promedio := float64(suma) / 3
	fmt.Println("La suma de las temperaturas es", suma)
	fmt.Println("El promedio de las temperaturas es", promedio)
}
