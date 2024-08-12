// Dado un número de 5 cifras, determinar si es capicúa.
// Si fuera un número de 6 cifras ¿Sirve la resolución planteada? ¿Cómo habría que modificarla?
package main

import (
	"fmt"
	"strconv"
)

func esCapicua(num int) bool {
	str := strconv.Itoa(num)

	return str[0] == str[4] && str[1] == str[3]
}

func main() {
	var num int

	fmt.Print("Ingresa un número de 5 cifras: ")
	fmt.Scan(&num)

	if esCapicua(num) {
		fmt.Println(num, "es capicúa")
	} else {
		fmt.Println(num, "no es capicúa")
	}
}
