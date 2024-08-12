// Ejercicio FizzBuzz en su forma clásica es el siguiente:
// "Para cada número del 1 al 100, imprime 'Fizz' si el número es divisible por 3, 'Buzz' si es divisible por 5 y 'FizzBuzz' si es divisible por ambos.
// Si no es divisible ni por 3 ni por 5, simplemente imprime el número."
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		output := ""

		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}

		fmt.Println(output)
	}
}
