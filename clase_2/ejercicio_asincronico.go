package main

import "fmt"

func main() {

	fmt.Println(ejercicio1(40000))
	fmt.Println(ejercicio1(80000))
	fmt.Println(ejercicio1(200000))

	fmt.Println(ejercicio2(4.5, 3.5, 5, 4))

}

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo.
// Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un
// salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 %
// del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
func ejercicio1(salario int32) float32 {

	var impuesto float32

	switch {
	case salario > 50000:
		impuesto = float32(salario) * 17 / 100
	case salario > 150000:
		impuesto = float32(salario) * 27 / 100
	default:
		impuesto = 0
	}

	return impuesto
}

// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita
// generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el
// promedio. No se pueden introducir notas negativas
func ejercicio2(notas ...float64) float64 {

	resultado := 0.0
	for _, value := range notas {
		resultado += value
	}

	return resultado / float64(len(notas))
}
