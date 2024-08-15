/*
1. Define una estructura llamada **`Persona`** con los siguientes campos: **`nombre`**, **`edad`**, **`ciudad`** y **`telefono`**.
2. Crea dos variables de tipo **`Persona`** con valores iniciales distintos.
3. Imprime los valores de ambas variables por pantalla.
4. Define una función llamada **`cambiarCiudad`** para cambiar la ciudad de la persona**.** Si la ciudad es la misma, no hace nada.
5. Llama a la función **`cambiarCiudad`** con una de las personas creadas anteriormente y una ciudad distinta a la actual.
6. Imprime los valores de ambas variables por pantalla para comprobar que se ha actualizado el campo **`ciudad`** de la persona correspondiente si la ciudad es distinta, o que no se ha actualizado si la ciudad es la misma.
*/
package main

import "fmt"

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

func main() {
	p1 := Persona{
		nombre:   "Juan",
		edad:     25,
		ciudad:   "La Plata",
		telefono: 22136974258,
	}
	p2 := Persona{
		nombre:   "Ana",
		edad:     31,
		ciudad:   "CABA",
		telefono: 1125754368,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.cambiarCiudad("El Calafate")

	fmt.Println(p1)
}

func (p *Persona) cambiarCiudad(city string) {
	if city != "" {
		p.ciudad = city
	}
}
