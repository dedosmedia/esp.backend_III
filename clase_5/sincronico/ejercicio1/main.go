package main

import (
	"fmt"
)

// Una empresa necesita realizar una buena gestión de sus emepleado, para esto realizaremos un pequeño
// programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

// Definiendo el método con el receiver
func (e Employee) PrintEmployee() {
	fmt.Printf("Employee: %v \n ", e)
}

func main() {
	fmt.Println("Hola")

	person1 := Person{
		ID:          1,
		Name:        "Jhon",
		DateOfBirth: "2013-01-01",
	}

	employee1 := Employee{
		ID:       2,
		Position: "Developer",
		Person:   person1,
	}

	employee1.PrintEmployee()
	//fmt.Printf()
}
