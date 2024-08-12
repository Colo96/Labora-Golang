package main

import "fmt"

func main() {
	fmt.Print("Ingrese una longitud: ")
	var length int
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Println("No se ingresó una longitud válida:", err)
		return
	}

	fmt.Print("Ingrese una anchura: ")
	var width int
	_, err = fmt.Scan(&width)
	if err != nil {
		fmt.Println("No se ingresó una anchura válida:", err)
		return
	}

	area := length * width
	perimeter := (length + width) * 2

	fmt.Println("Area:", area, "Perímetro:", perimeter)
}
