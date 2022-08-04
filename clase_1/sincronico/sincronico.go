package main

import (
	"fmt"
)

// Antes de comenzar el main
func init() {
	fmt.Println("Antes del programa principal")
}

func main() {

	// Inicialización de arreglos, diferentes formas

	var arreglos [5]int
	fmt.Printf("El valor es %v y el tipo %T \n", arreglos, arreglos)

	vals := []int{1, 2, 3, 4, 5}
	fmt.Printf("El valor es %v y el tipo %T \n", vals, vals)

	var y [5]int = [5]int{10, 11, 12}
	fmt.Printf("El valor es %v y el tipo %T \n", y, y)

	x := [...]int{10, 11, 12}
	fmt.Printf("El valor es %v y el tipo %T \n", x, x)

	z := [5]int{1: 10, 4: 12}
	fmt.Printf("El valor es %v y el tipo %T \n", z, z)

	// Slices

	var slice1 = make([]int, 10) // solo longitud
	fmt.Printf("%v Tipo %T len: %v  cap: %v", slice1, slice1, len(slice1), cap(slice1))
	fmt.Println("")

	var slice2 = make([]int, 3, 20) // longitud y capacidad
	fmt.Printf("%v Tipo %T len: %v  cap: %v", slice2, slice2, len(slice2), cap(slice2))
	fmt.Println("")

	var window = z[0:1]
	fmt.Printf("%v Tipo %T len: %v  cap: %v", window, window, len(window), cap(window))
	fmt.Println("")

	// Maps  (clave:valor)

	// explicito
	var employee = map[string]int{"mark": 1, "diego": 2}
	fmt.Printf("%v Tipo %T len: %v  ", employee, employee, len(employee))
	fmt.Println("")

	// con make
	var employee2 = make(map[string]int)
	fmt.Printf("%v Tipo %T len: %v  ", employee2, employee2, len(employee2))
	fmt.Println("")

	// implicito
	paises := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("%v Tipo %T len: %v ", paises, paises, len(paises))
	fmt.Println("")

}
