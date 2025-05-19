package product

import (
	"context"
	"errors"
	"github.com/y4ns3/sarkor-test/internal/entity"
)

type Service struct {
	repo ProductRepository
}

func NewService(p ProductRepository) *Service {
	return &Service{repo: p}
}
func (s *Service) CreateProduct(ctx context.Context, name, description string, price int64, quantity int) (int, error) {
	if quantity < 0 {
		return -1, ErrProductQuantityCannotBeNegative
	}
	if price < 0 {
		return -1, ErrPriceCannotBeNegative
	}
	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	id, err := s.repo.CreateProduct(ctx, &product)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (s *Service) GetProductByID(ctx context.Context, id int) (*entity.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}
func (s *Service) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *Service) UpdateProduct(ctx context.Context, ID int, name, description string, quantity int, price int64) error {
	product, err := s.repo.GetProductByID(ctx, ID)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrProductNotFound
	}
	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if quantity != 0 {
		product.Quantity = quantity
	}
	if price != 0 {
		product.Price = price
	}
	err = s.repo.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteProduct(ctx context.Context, id int) error {
	err := s.repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

var (
	ErrProductNotFound                 = errors.New("product is not found")
	ErrProductQuantityCannotBeNegative = errors.New("product quantity cannot be negative")
	ErrPriceCannotBeNegative           = errors.New("product price cannot be negative")
)
