package main

import (
	"errors"
	"fmt"
)

const (
	tipoPequeno = "pequeño"
	tipoMediano = "mediano"
	tipoGrande  = "grande"
)

type IProducto interface {
	Precio() float32
}

type Producto struct {
	costo float32
}

func (p Producto) Precio() float32 {
	return 0
}

type Pequeno struct {
	producto Producto
}

type Mediano struct {
	producto Producto
}

type Grande struct {
	producto Producto
}

func (p Pequeno) Precio() float32 {
	return p.producto.costo
}
func (p Mediano) Precio() float32 {
	return p.producto.costo + p.producto.costo*3.0/100.0
}
func (p Grande) Precio() float32 {
	return p.producto.costo + p.producto.costo*3.0/100.0 + 2500.0
}

func factoryProduct(tipo string, precio float32) (IProducto, error) {

	switch tipo {
	case tipoPequeno:
		return Pequeno{
			producto: Producto{
				costo: precio,
			},
		}, nil
	case tipoMediano:
		return Mediano{
			producto: Producto{
				costo: precio,
			},
		}, nil
	case tipoGrande:
		return Grande{
			producto: Producto{
				costo: precio,
			},
		}, nil
	default:
		return Producto{}, errors.New("No existe producto")
	}

}

func main() {

	p, err := factoryProduct(tipoPequeno, 50.0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(p.Precio())

	p, err = factoryProduct(tipoMediano, 50.0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(p.Precio())

	p, err = factoryProduct(tipoGrande, 50.0)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(p.Precio())

	p, err = factoryProduct("otro", 50.0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
