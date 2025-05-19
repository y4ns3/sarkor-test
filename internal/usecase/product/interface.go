package product

import (
	"context"
	"github.com/y4ns3/sarkor-test/internal/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *entity.Product) (int, error)
	GetProductByID(ctx context.Context, id int) (*entity.Product, error)
	GetProducts(ctx context.Context) ([]*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id int) error
}
