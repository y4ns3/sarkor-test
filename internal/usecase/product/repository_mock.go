package product

import (
	"context"
	"errors"
	"github.com/y4ns3/sarkor-test/internal/entity"
)

var products = []*entity.Product{
	{ID: 1, Name: "name", Description: "description1", Price: 2001_00, Quantity: 1},
	{ID: 2, Name: "name1", Description: "description2", Price: 2002_00, Quantity: 3},
	{ID: 3, Name: "name2", Description: "description3", Price: 2023_00, Quantity: 5},
	{ID: 4, Name: "name3", Description: "description4", Price: 2004_00, Quantity: 6},
}
var lastID = 4

type MockProductRepository struct {
	products []*entity.Product
}

func NewMockRepository() *MockProductRepository {
	return &MockProductRepository{products: products}
}

func (m *MockProductRepository) CreateProduct(_ context.Context, product *entity.Product) (int, error) {
	lastID++
	newProduct := entity.Product{
		ID:          lastID,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Name:        product.Name,
	}
	m.products = append(m.products, &newProduct)
	return lastID, nil
}
func (m *MockProductRepository) GetProductByID(_ context.Context, id int) (*entity.Product, error) {
	for _, product := range m.products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}
func (m *MockProductRepository) GetProducts(_ context.Context) ([]*entity.Product, error) {
	return m.products, nil
}

func (m *MockProductRepository) UpdateProduct(_ context.Context, product *entity.Product) error {
	if product == nil {
		return errors.New("product is nil")
	}
	for _, p := range m.products {
		if p.ID == product.ID {
			p.Name = product.Name
			p.Price = product.Price
			p.Quantity = product.Quantity
			p.Description = product.Description
			return nil
		}
	}
	return errors.New("product not found")
}
func (m *MockProductRepository) DeleteProduct(_ context.Context, id int) error {
	for i, p := range m.products {
		if p.ID == id {
			m.products = append(m.products[:i], m.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
