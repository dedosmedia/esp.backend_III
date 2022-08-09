package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{1, "A", 10.0, "Desc1", "Cat1"},
	{2, "B", 11.0, "Desc2", "Cat2"},
	{3, "C", 12.0, "Desc3", "Cat3"},
}

func (p Product) Save() {
	Products = append(Products, p)
}
func (p Product) GetAll() {

	for k, v := range Products {
		fmt.Printf("[%v]: %v \n", k, v)
	}
}

func getById(i int) Product {

	for _, v := range Products {
		if v.ID == i {
			return v
		}
	}

	return Product{}
}

func main() {

	p := Product{4, "D", 13.0, "Desc4", "Cat4"}
	p.GetAll()

	p.Save()
	p.GetAll()

	fmt.Printf("%v \n", getById(2))

}
