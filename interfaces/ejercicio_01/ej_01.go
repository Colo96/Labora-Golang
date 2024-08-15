package main

import (
	"fmt"
	"math/rand"
)

// Generator es la interfaz que define cómo se generan los números.
type Generator interface {
	Next() int
}

// Fulfill es la interfaz que define el método para verificar si un número cumple una condición.
type Fulfill interface {
	Fulfill(n int) bool
}

// IncreasingGenerator genera números en orden creciente.
type IncreasingGenerator struct {
	current int
}

// Next retorna el siguiente número en orden creciente.
func (gen *IncreasingGenerator) Next() int {
	gen.current++
	return gen.current
}

// DecreasingGenerator genera números en orden decreciente.
type DecreasingGenerator struct {
	current int
}

// Next retorna el siguiente número en orden decreciente.
func (gen *DecreasingGenerator) Next() int {
	gen.current--
	return gen.current
}

// Sequence combina un generador y un predicado.
type Sequence struct {
	generator Generator
	predicate Fulfill
}

// Next retorna el próximo número que cumple con el predicado.
func (seq *Sequence) Next() int {
	for {
		num := seq.generator.Next()
		if seq.predicate.Fulfill(num) {
			return num
		}
	}
}

// AnyNumber cumple con la condición de que cualquier número es válido.
type AnyNumber struct{}

// Fulfill siempre retorna true.
func (a AnyNumber) Fulfill(n int) bool {
	return true
}

// PrimeNumber cumple con la condición de ser un número primo.
type PrimeNumber struct{}

// Fulfill retorna true si el número es primo.
func (p PrimeNumber) Fulfill(n int) bool {
	return isPrime(n)
}

// EvenNumber cumple con la condición de ser un número par.
type EvenNumber struct{}

// Fulfill retorna true si el número es par.
func (e EvenNumber) Fulfill(n int) bool {
	return n%2 == 0
}

// isPrime verifica si un número es primo.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Elegir aleatoriamente entre un generador creciente o decreciente
	var generator Generator
	if rand.Intn(2) == 0 {
		fmt.Println("Generando en orden creciente:")
		generator = &IncreasingGenerator{current: 0}
	} else {
		fmt.Println("Generando en orden decreciente:")
		generator = &DecreasingGenerator{current: 100} // Comienza en 100 para decreciente
	}

	// Elegir aleatoriamente un predicado
	var predicate Fulfill
	switch rand.Intn(3) {
	case 0:
		fmt.Println("Cualquier número:")
		predicate = AnyNumber{}
	case 1:
		fmt.Println("Números primos:")
		predicate = PrimeNumber{}
	case 2:
		fmt.Println("Números pares:")
		predicate = EvenNumber{}
	}

	// Crear la secuencia
	sequence := Sequence{generator: generator, predicate: predicate}

	// Imprimir los primeros 30 números de la secuencia
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", sequence.Next())
	}
	fmt.Println()
}
