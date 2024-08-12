// Redactar un algoritmo para pasar un periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.
package main

import "fmt"

func main() {
	dias := 32
	horas := 7
	minutos := 55
	segundos := 3600

	fmt.Println("Los campos", dias, "dias:", horas, "horas:", minutos, "minutos:", segundos, "segundos son", calcSec(dias, horas, minutos, segundos), "segundos totales")
}

func calcSec(dias, horas, minutos, segundos int) int {
	return (((dias * 24) * 60) * 60) + ((horas * 60) * 60) + (minutos * 60) + segundos
}
