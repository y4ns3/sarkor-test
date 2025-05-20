package handler

import (
	"context"

	"github.com/y4ns3/sarkor-test/internal/entity"
)

type ProductService interface {
	CreateProduct(ctx context.Context, name, description string, price int64, quantity int) (int, error)
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
	GetProducts(ctx context.Context) ([]*entity.Product, error)
	UpdateProduct(ctx context.Context, ID int, name, description string, quantity int, price int64) error
	DeleteProduct(ctx context.Context, id int) error
}
