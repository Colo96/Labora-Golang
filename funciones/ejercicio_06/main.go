// Realice un algoritmo que dado un string te diga si es palindromo
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ingresarStr(mensaje string) (string, error) {
	fmt.Print(mensaje)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		cadena := strings.TrimSpace(scanner.Text())
		if cadena == "" {
			return "", fmt.Errorf("la entrada no puede estar vacía")
		}
		return cadena, nil
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error al leer la entrada: %v", err)
	}
	return "", fmt.Errorf("no se pudo leer la entrada")
}

func esPalindromo(cadena string) bool {
	runes := []rune(cadena)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) == cadena
}

func main() {
	cadena, err := ingresarStr("Ingrese una cadena: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if esPalindromo(cadena) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
