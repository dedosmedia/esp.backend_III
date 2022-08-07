package main

import "fmt"

/*
Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos. Para ello es necesario generar una estructura Alumno con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle.
*/

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {

	estudiante1 := Estudiante{
		Nombre:   "Diego",
		Apellido: "Díaz",
		DNI:      616626,
		Fecha:    "12/04/2000",
	}

	fmt.Printf("El señor: %s %s \nDNI: %v\nNació: %s", estudiante1.Nombre, estudiante1.Apellido, estudiante1.DNI, estudiante1.Fecha)

}
