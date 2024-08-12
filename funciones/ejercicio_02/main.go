// Realice un algoritmo para determinar si dos números son amigos.
// Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
package main

import "fmt"

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

func sumaDivisores(num int) int {
	suma := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			suma += i
		}
	}
	return suma
}

func sonAmigos(num1, num2 int) bool {
	return sumaDivisores(num1) == num2 && sumaDivisores(num2) == num1
}

func main() {
	num1, err := pedirNumero("Ingrese un número: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	num2, err := pedirNumero("Ingrese un segundo número: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if sonAmigos(num1, num2) {
		fmt.Println(num1, "es amigo de", num2)
	} else {
		fmt.Println(num1, "no es amigo de", num2)
	}
}
