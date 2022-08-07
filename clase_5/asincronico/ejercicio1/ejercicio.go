package main

import (
	"fmt"
)

type Autor struct {
	Nombre   string
	Apellido string
}

func (a Autor) nombreCompleto() string {

	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo      string
	Descripcion string
	Autor       Autor
}

func (l Libro) informacion() {
	fmt.Println("Título", l.Titulo)
	fmt.Println("Descripcion", l.Descripcion)
	fmt.Println("Autor", l.Autor.nombreCompleto())
}

func main() {

	autor := Autor{
		Nombre:   "Diego",
		Apellido: "Diaz",
	}

	libro := Libro{
		Titulo:      "El libro",
		Descripcion: "La descripcion",
		Autor:       autor,
	}

	libro.informacion()

}
