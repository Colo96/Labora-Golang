// Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
package main

import "fmt"

func main() {
	segundosEjemplos := []int{1030, 12045, 176520}

	for _, segundos := range segundosEjemplos {
		dias, horas, minutos, segundosRestantes := convertirSegundos(segundos)
		fmt.Printf("%d segundos son %d días, %d horas, %d minutos con %d segundos\n", segundos, dias, horas, minutos, segundosRestantes)
	}
}

func convertirSegundos(segundosTotales int) (dias int, horas int, minutos int, segundos int) {
	dias = segundosTotales / 86400
	segundosTotales %= 86400
	horas = segundosTotales / 3600
	segundosTotales %= 3600
	minutos = segundosTotales / 60
	segundos = segundosTotales % 60
	return
}
