// Realice un algoritmo que dado un número te diga si es capicua.
package main

import (
	"fmt"
	"strconv"
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

func esCapicua(num int) bool {
	strNum := strconv.Itoa(num)
	runes := []rune(strNum)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return strNum == string(runes)
}

func main() {
	num, err := pedirNumero("Ingrese un número: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if esCapicua(num) {
		fmt.Println("Es capicua")
	} else {
		fmt.Println("No es capicua")
	}
}
