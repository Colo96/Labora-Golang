// Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.
package main

import "fmt"

func main() {
	X := 860
	S1, S2, S3, S4, S5 := descomponerNumero(X)
	fmt.Printf("X = %d se descompone en: S1 = %d, S2 = %d, S3 = %d, S4 = %d, S5 = %d\n", X, S1, S2, S3, S4, S5)
}

func descomponerNumero(X int) (int, int, int, int, int) {
	var S1, S2, S3, S4, S5 int

	if X > 50 {
		S1 = 50
		X -= 50
	} else {
		S1 = X
		X = 0
	}

	if X > 50 {
		S2 = 50
		X -= 50
	} else {
		S2 = X
		X = 0
	}

	if X > 600 {
		S3 = 600
		X -= 600
	} else {
		S3 = X
		X = 0
	}

	if X > 800 {
		S4 = 800
		X -= 800
	} else {
		S4 = X
		X = 0
	}

	S5 = X

	return S1, S2, S3, S4, S5
}
