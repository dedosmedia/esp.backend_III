package main

import (
	"fmt"
)

func main() {

	ejercicio1(200, 16)

	ejercicio2(30, 25, 150000)
	ejercicio2(18, 25, 150000)
	ejercicio2(30, 8, 150000)
	ejercicio2(30, 25, 50000)

}

// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre
// sus productos. Para ello necesitan una aplicación que les permita
// calcular el descuento basándose en dos variables: su precio y el
// descuento en porcentaje. La tienda espera obtener como resultado
// el valor con el descuento aplicado y luego imprimirlo en consola.

func ejercicio1(precio float32, descuento int8) {

	total := precio - precio*float32(descuento)/100

	fmt.Println("El valor con descuento es ", total)

}

// Un banco quiere otorgar préstamos a sus clientes, pero no todos
// pueden acceder a los mismos. El banco tiene ciertas reglas para
// saber a qué cliente se le puede otorgar: solo le otorga préstamos a
// clientes cuya edad sea mayor a 22 años, se encuentren empleados y
// tengan más de un año de antigüedad en su trabajo. Dentro de los
// préstamos que otorga, no les cobrará interés a los que su sueldo
// sea mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y
// que imprima un mensaje de acuerdo a cada caso
func ejercicio2(edad int8, antiguedadMeses int16, sueldo int32) {

	if edad > 22 && antiguedadMeses > 12 {
		fmt.Println("Sí te otorgaremos un préstamo")
		if sueldo > 100000 {
			fmt.Println("Y sin cobrar tasa de interés")
		} else {
			fmt.Println("pero con cobro de tasa de interés")
		}
	} else {
		fmt.Println("Lo siento no te podemos prestar")
	}

}
