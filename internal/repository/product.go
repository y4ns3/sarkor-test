package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/y4ns3/sarkor-test/internal/entity"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
func (p *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) (int, error) {
	query := `INSERT INTO products(name,description,price,quantity) VALUES ($1,$2,$3,$4) RETURNING id`
	var id int
	err := p.db.QueryRow(ctx, query, product.Name, product.Description, product.Price, product.Quantity).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil

}

func (p *ProductRepository) GetProductByID(ctx context.Context, id int) (*entity.Product, error) {
	query := `SELECT id,name,description,price,quantity FROM products WHERE id = $1`
	row := p.db.QueryRow(ctx, query, id)
	product := &entity.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}
func (p *ProductRepository) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	query := `SELECT id,name,description,price,quantity FROM products`
	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, ErrProductsNotFound
	}
	return products, nil
}
func (p *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id = $5`
	_, err := p.db.Exec(ctx, query, product.Name, product.Description, product.Price, product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrProductsNotFound = errors.New("products not found")
)
