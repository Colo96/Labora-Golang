// Redactar un algoritmo que pida al usuario el ingreso de una serie de números reales e imprima por pantalla el resultado de sumarlos.
// La serie finaliza cuando el usuario ingresa el número cero.
package main

import (
	"fmt"
	"log"
)

func main() {
	var numero int
	for {
		fmt.Print("Ingrese un número: ")
		_, err := fmt.Scan(&numero)
		if err != nil {
			log.Printf("Error al leer la entrada: %v. Por favor, intente de nuevo.", err)
			continue
		}
		if numero < 1 {
			fmt.Println("Número menor que 1 ingresado. Terminando el programa.")
			break
		}
		fmt.Printf("El usuario ingresó %d y la suma es %d\n", numero, sumarNumeros(numero))
	}
}

func sumarNumeros(numero int) int {
	return numero * (numero + 1) / 2
}
