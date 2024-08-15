/*
Defina una estructura de datos para manejar conjuntos de enteros de la siguiente forma.

	`Add(value int) int`

Agrega al final. Se incrementa en uno la longitud.

`RemoveAt(index int) bool`

True si lo borró, false si no pudo (no importa el motivo). Si se borra entonces se decrementa en uno la longitud

`Lenght() int`

Longitud: Cantidad de elementos que tiene. Esta determinada por la cantidad de valores que se agregan menos aquellos que se borran.

`Set(value int, index int) bool`

True si lo modificó, false si no pudo (no importa el motivo)

`ToString() string`

Libre elección. La función strings.Join (del paquete strings) puede ser de útil.

Para pensar: Tanto el `RemoveAt` y el `Set` son métodos cuya ejecución puede “fallar” si se pasa como `index` un valor que esté fuera de los “límites” (dado por la longitud)… Yo pregunto entonces 🤔 ¿Es esto algo entonces que pueda ser codificado en una función aparte de forma tal que se re utilice en ambos métodos?

`isInBoundaries(index) bool`
*/
package main

import (
	"fmt"
	"strings"
)

// IntegerSet representa un conjunto de enteros.
type IntegerSet struct {
	elements []int
}

// isInBoundaries verifica si el índice dado está dentro de los límites del conjunto.
func (set *IntegerSet) isInBoundaries(index int) bool {
	return index >= 0 && index < len(set.elements)
}

// Add agrega un valor al final del conjunto.
func (set *IntegerSet) Add(value int) int {
	set.elements = append(set.elements, value)
	return len(set.elements)
}

// RemoveAt elimina el elemento en el índice dado.
// Retorna true si el elemento fue eliminado, false si no.
func (set *IntegerSet) RemoveAt(index int) bool {
	if !set.isInBoundaries(index) {
		return false
	}
	set.elements = append(set.elements[:index], set.elements[index+1:]...)
	return true
}

// Length retorna la cantidad de elementos en el conjunto.
func (set *IntegerSet) Length() int {
	return len(set.elements)
}

// Set modifica el valor en el índice dado.
// Retorna true si el valor fue modificado, false si no.
func (set *IntegerSet) Set(value int, index int) bool {
	if !set.isInBoundaries(index) {
		return false
	}
	set.elements[index] = value
	return true
}

// ToString convierte el conjunto de enteros a una cadena de texto.
func (set *IntegerSet) ToString() string {
	strValues := make([]string, len(set.elements))
	for i, v := range set.elements {
		strValues[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(strValues, ", ")
}

func main() {
	set := &IntegerSet{}

	set.Add(10)
	set.Add(20)
	set.Add(30)

	fmt.Println("Conjunto:", set.ToString())
	fmt.Println("Longitud:", set.Length())

	set.Set(25, 1)
	fmt.Println("Conjunto después de modificar:", set.ToString())

	set.RemoveAt(1)
	fmt.Println("Conjunto después de remover:", set.ToString())
	fmt.Println("Longitud después de remover:", set.Length())
}
