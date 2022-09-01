package product

import (
	"actividad/clase/internal/domain"
	"errors"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	DeleteByID(id int) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los productos
func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca un producto por id
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
func (s *service) Update(p domain.Product) (domain.Product, error) {
	p, err := s.r.Update(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// DeleteByID elimina un producto por id
func (s *service) DeleteByID(id int) error {
	err := s.r.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
