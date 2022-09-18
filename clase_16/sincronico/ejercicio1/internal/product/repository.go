package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	Patch(p domain.Product) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

type repository struct {
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// Update actualiza un nuevo producto
func (r *repository) Update(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}

	for k, v := range r.list {
		if v.Id == p.Id {
			r.list[k] = p
		}
	}
	return p, nil
}

// Patch patchea un viejo producto
func (r *repository) Delete(id int) (domain.Product, error) {
	_, err := r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}

	for k, v := range r.list {

		if v.Id == id {

			r.list = append(r.list[:k], r.list[k+1:]...)
			return v, nil
		}
	}
	return domain.Product{}, errors.New("Nunca debí entrar aquí")

}

// Patch patchea un viejo producto
func (r *repository) Patch(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}

	p2, err := r.GetByID(p.Id)
	if err != nil {
		return domain.Product{}, err
	}

	if p.Name != "" {
		p2.Name = p.Name
	}
	if p.Quantity != 0 {
		p2.Quantity = p.Quantity
	}
	if p.CodeValue != "" {
		p2.CodeValue = p.CodeValue
	}
	/*
		if p.IsPublished != false {
			p2.Name = p.Name
		}
	*/
	if p.Expiration != "" {
		p2.Expiration = p.Expiration
	}
	if p.Price != 0 {
		p2.Price = p.Price
	}

	for k, v := range r.list {
		if v.Id == p.Id {
			r.list[k] = p2
		}
	}
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}
