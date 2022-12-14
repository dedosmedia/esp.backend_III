package product

import (
	"errors"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Patch(id int, p domain.Product) (domain.Product, error)
	Delete(id int) (domain.Product, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos} los productos
func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca un producto por su id
func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// SearchPriceGt busca productos por precio mayor que el precio dado
func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return l, nil
}

// Create agrega un nuevo producto
func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// Update actualiza un nuevo producto
func (s *service) Update(id int, p domain.Product) (domain.Product, error) {
	p.Id = id

	p, err := s.r.Update(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// Patch patchea un viejo producto
func (s *service) Patch(id int, p domain.Product) (domain.Product, error) {
	p.Id = id

	p, err := s.r.Patch(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil

}

// Delete elimina  un viejo producto
func (s *service) Delete(id int) (domain.Product, error) {

	p, err := s.r.Delete(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil

}
