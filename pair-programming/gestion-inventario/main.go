package main

import (
	"errors"
	"fmt"
	"strings"
)

// InventarioOperaciones define la interfaz para las operaciones que se pueden realizar en un inventario
type InventarioOperaciones interface {
	AgregarProducto(prod Producto) error                                       // Agrega un producto al inventario
	EliminarProducto(nombre string) error                                      // Elimina un producto del inventario por su nombre
	MostrarProductos() error                                                   // Muestra todos los productos del inventario
	ModificarCantidad(nombre string, nuevaCant int) error                      // Modifica la cantidad de un producto existente
	ModificarProducto(nombre string, nuevoPrecio float64, nuevaCant int) error // Modifica el precio y la cantidad de un producto existente
}

// Inventario es una estructura que implementa la interfaz InventarioOperaciones
type Inventario struct {
	productos []Producto // Lista de productos en el inventario
}

// Producto representa un producto con su nombre, precio y cantidad
type Producto struct {
	nombre   string  // Nombre del producto
	precio   float64 // Precio del producto
	cantidad int     // Cantidad disponible del producto
}

// AgregarProducto agrega un nuevo producto al inventario
// Retorna un error si el precio o la cantidad son negativos
func (inv *Inventario) AgregarProducto(prod Producto) error {
	if prod.precio < 0 {
		return errors.New("el precio no puede ser negativo")
	}
	if prod.cantidad < 0 {
		return errors.New("la cantidad no puede ser negativa")
	}
	inv.productos = append(inv.productos, prod) // Añade el producto al inventario
	return nil
}

// EliminarProducto elimina un producto del inventario basado en su nombre
// Retorna un error si el producto no se encuentra en el inventario
func (inv *Inventario) EliminarProducto(nombre string) error {
	for i, v := range inv.productos {
		// Compara el nombre de manera insensible a mayúsculas/minúsculas
		if strings.EqualFold(nombre, v.nombre) {
			// Elimina el producto de la lista
			inv.productos = append(inv.productos[:i], inv.productos[i+1:]...)
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

// MostrarProductos muestra todos los productos en el inventario
// Retorna un error si el inventario está vacío
func (inv *Inventario) MostrarProductos() error {
	if len(inv.productos) == 0 {
		return errors.New("el inventario está vacío")
	}
	fmt.Println("Inventario:")
	for _, prod := range inv.productos {
		fmt.Printf("Nombre: %s | Precio: $%.2f | Cantidad: %d\n", prod.nombre, prod.precio, prod.cantidad)
	}
	return nil
}

// ModificarCantidad modifica la cantidad de un producto existente en el inventario
// Retorna un error si la nueva cantidad es negativa o si el producto no se encuentra
func (inv *Inventario) ModificarCantidad(nombre string, nuevaCant int) error {
	if nuevaCant < 0 {
		return errors.New("la cantidad no puede ser negativa")
	}
	for i, v := range inv.productos {
		// Busca el producto por nombre
		if strings.EqualFold(nombre, v.nombre) {
			inv.productos[i].cantidad = nuevaCant // Actualiza la cantidad
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

// ModificarProducto modifica tanto el precio como la cantidad de un producto existente
// Retorna un error si el nuevo precio o la nueva cantidad son negativos, o si el producto no se encuentra
func (inv *Inventario) ModificarProducto(nombre string, nuevoPrecio float64, nuevaCant int) error {
	if nuevoPrecio < 0 {
		return errors.New("el precio no puede ser negativo")
	}
	if nuevaCant < 0 {
		return errors.New("la cantidad no puede ser negativa")
	}
	for i, v := range inv.productos {
		// Busca el producto por nombre
		if strings.EqualFold(nombre, v.nombre) {
			// Actualiza el precio y la cantidad
			inv.productos[i].precio = nuevoPrecio
			inv.productos[i].cantidad = nuevaCant
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

// main es el punto de entrada del programa
// Aquí se implementa el ciclo de interacción con el usuario
func main() {
	// Creamos un inventario que implementa la interfaz InventarioOperaciones
	var inv InventarioOperaciones = &Inventario{}
	var nombre string
	var precio float64
	var cant int

	for {
		// Muestra las opciones disponibles al usuario
		fmt.Println("Seleccione una opción: Agregar/Eliminar/Mostrar/ModCant/ModProd/Salir")
		var op string
		fmt.Scan(&op) // Lee la opción ingresada por el usuario

		switch strings.ToLower(op) { // Convierte la opción a minúsculas para evitar problemas con la capitalización
		case "agregar":
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingrese el precio del producto: ")
			fmt.Scan(&precio)
			fmt.Print("Ingrese la cantidad del producto: ")
			fmt.Scan(&cant)

			// Crea un nuevo producto y lo agrega al inventario
			prod := Producto{
				nombre:   nombre,
				precio:   precio,
				cantidad: cant,
			}
			// Maneja el posible error al agregar el producto
			if err := inv.AgregarProducto(prod); err != nil {
				fmt.Println("Error al agregar producto:", err)
			} else {
				fmt.Println("Producto agregado con éxito.")
			}
		case "eliminar":
			fmt.Print("Ingrese el nombre del producto que quiere eliminar: ")
			fmt.Scan(&nombre)
			// Maneja el posible error al eliminar el producto
			if err := inv.EliminarProducto(nombre); err != nil {
				fmt.Println("Error al eliminar producto:", err)
			} else {
				fmt.Println("Producto eliminado con éxito.")
			}
		case "mostrar":
			// Maneja el posible error al mostrar los productos
			if err := inv.MostrarProductos(); err != nil {
				fmt.Println("Error al mostrar productos:", err)
			}
		case "modcant":
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingrese la nueva cantidad: ")
			fmt.Scan(&cant)
			// Maneja el posible error al modificar la cantidad
			if err := inv.ModificarCantidad(nombre, cant); err != nil {
				fmt.Println("Error al modificar la cantidad:", err)
			} else {
				fmt.Println("Cantidad modificada con éxito.")
			}
		case "modprod":
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingrese el nuevo precio: ")
			fmt.Scan(&precio)
			fmt.Print("Ingrese la nueva cantidad: ")
			fmt.Scan(&cant)
			// Maneja el posible error al modificar el producto
			if err := inv.ModificarProducto(nombre, precio, cant); err != nil {
				fmt.Println("Error al modificar el producto:", err)
			} else {
				fmt.Println("Producto modificado con éxito.")
			}
		case "salir":
			// Termina el ciclo y sale del programa
			return
		default:
			// Maneja el caso donde se ingresa una opción inválida
			fmt.Println("Opción no válida.")
		}
	}
}
