package main

import (
	"fmt"
)

/*
Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña. Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.
*/

type Usuarios struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (u *Usuarios) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuarios) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuarios) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuarios) cambiarContrasena(contrasena string) {
	u.Contrasena = contrasena
}
func (u Usuarios) imprimir() {

	fmt.Println("Nombre: ", u.Nombre)
	fmt.Println("Apellido: ", u.Apellido)
	fmt.Println("Edad: ", u.Edad)
	fmt.Println("Correo: ", u.Correo)
	fmt.Println("Contrasena: ", u.Contrasena)
}

func main() {

	u1 := Usuarios{
		Nombre:     "Diego",
		Apellido:   "Díaz",
		Edad:       40,
		Correo:     "diego@dedosmedia.com",
		Contrasena: "12345",
	}

	u1.imprimir()

	u1.cambiarNombre("Pedro", "Pérez")

	u1.imprimir()

}
