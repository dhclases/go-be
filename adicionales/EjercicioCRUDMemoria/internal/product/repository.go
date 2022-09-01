package product

import (
	"actividad/clase/internal/domain"
	"errors"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(p domain.Product) (domain.Product, error)
	DeleteByID(id int) error
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
func (r *repository) Update(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValueExceptId(p.CodeValue, p.Id) {
		return domain.Product{}, errors.New("code value already exists")
	}
	originalProduct, err := r.GetByID(p.Id)
	if err != nil {
		return p, err
	}
	index, err := r.getIndex(p.Id)
	if err != nil {
		return domain.Product{}, err
	}
	r.list[index] = setProduct(p, originalProduct)
	return r.list[index], nil
}

// Update actualiza un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}

	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

func (r *repository) DeleteByID(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	index, err := r.getIndex(id)
	if err != nil {
		return err
	}
	r.list = removeItem(r.list, index)
	return nil
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

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValueExceptId(codeValue string, id int) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue && product.Id != id {
			return false
		}
	}
	return true
}

func setProduct(new domain.Product, original domain.Product) domain.Product {
	original.CodeValue = new.CodeValue
	original.Expiration = new.Expiration
	original.IsPublished = new.IsPublished
	original.Name = new.Name
	original.Price = new.Price
	original.Quantity = new.Quantity
	return original
}

func (r *repository) getIndex(id int) (int, error) {
	for i, product := range r.list {
		if product.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("no such product")
}

func removeItem(s []domain.Product, i int) []domain.Product {

	return append(s[:i], s[i+1:]...)

}
