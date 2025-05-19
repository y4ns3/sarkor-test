package product

import (
	"context"
	"testing"

	"github.com/y4ns3/sarkor-test/internal/entity"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := NewMockRepository()
	productService := NewService(mockRepo)
	ctx := context.Background()

	newProduct := &entity.Product{
		Name:        "new product",
		Description: "test",
		Price:       123321_00,
		Quantity:    2,
	}

	expectedID := 5
	id, err := productService.CreateProduct(ctx, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity)
	if err != nil {
		t.Fatalf("failed to create product: %v", err)
	}
	if id != expectedID {
		t.Errorf("expected ID %d, got %d", expectedID, id)
	}
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := NewMockRepository()
	productService := NewService(mockRepo)
	ctx := context.Background()

	newProduct := &entity.Product{
		Name:        "updated name",
		Description: "updated description",
		Price:       500000,
		Quantity:    10,
	}

	t.Run("update non-existent product", func(t *testing.T) {
		err := productService.UpdateProduct(ctx, 9999, newProduct.Name, newProduct.Description, newProduct.Quantity, newProduct.Price)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
	})

	t.Run("update existing product", func(t *testing.T) {
		err := productService.UpdateProduct(ctx, 1, newProduct.Name, newProduct.Description, newProduct.Quantity, newProduct.Price)
		if err != nil {
			t.Fatalf("failed to update existing product: %v", err)
		}

		updatedProduct, err := productService.GetProductByID(ctx, 1)
		if err != nil {
			t.Fatalf("failed to fetch updated product: %v", err)
		}
		if updatedProduct.Name != newProduct.Name || updatedProduct.Description != newProduct.Description {
			t.Error("product was not updated correctly")
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := NewMockRepository()
	productService := NewService(mockRepo)
	ctx := context.Background()

	t.Run("delete existing product", func(t *testing.T) {
		err := productService.DeleteProduct(ctx, 1)
		if err != nil {
			t.Fatalf("failed to delete product: %v", err)
		}
	})

	t.Run("delete non-existent product", func(t *testing.T) {
		err := productService.DeleteProduct(ctx, 333333)
		if err == nil {
			t.Error("expected error when deleting non-existent product, got nil")
		}
	})
}
func TestGetProductByID(t *testing.T) {
	mockRepo := NewMockRepository()
	productService := NewService(mockRepo)
	ctx := context.Background()

	t.Run("existing product", func(t *testing.T) {
		product, err := productService.GetProductByID(ctx, 4)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if product == nil {
			t.Fatal("expected product, got nil")
		}
		if product.ID != 4 {
			t.Errorf("expected ID 1, got %d", product.ID)
		}
	})

	t.Run("non-existent product", func(t *testing.T) {
		product, err := productService.GetProductByID(ctx, 9999)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if product != nil {
			t.Errorf("expected nil product, got %+v", product)
		}
	})
}

func TestGetProducts(t *testing.T) {
	ctx := context.Background()
	mockRepo := NewMockRepository()
	productService := NewService(mockRepo)

	t.Run("returns all existing products", func(t *testing.T) {
		expectedCount := len(mockRepo.products)

		products, err := productService.GetProducts(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(products) != expectedCount {
			t.Errorf("expected %d products, got %d", expectedCount, len(products))
		}
	})

	t.Run("returns empty slice when no products", func(t *testing.T) {
		mockRepo.products = []*entity.Product{}

		products, err := productService.GetProducts(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(products) != 0 {
			t.Errorf("expected 0 products, got %d", len(products))
		}
	})
}
