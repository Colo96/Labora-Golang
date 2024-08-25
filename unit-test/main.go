package main

import (
	"fmt"
	"unit-test/test"
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

func main() {
	num, err := pedirNumero("Ingrese un número: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := test.SumOddNumbers(num)
	fmt.Println("La suma de los primeros numeros impares es:", result)
}
