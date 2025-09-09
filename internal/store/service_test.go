package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStore(t *testing.T) {
	svc := NewService()
	ctx := context.Background()

	store, err := svc.CreateStore(ctx, "1", "Test Store")
	assert.NoError(t, err)
	assert.Equal(t, "Test Store", store.Name)

	_, err = svc.CreateStore(ctx, "1", "Duplicate Store")
	assert.Error(t, err)
}

func TestAddProduct(t *testing.T) {
	svc := NewService()
	ctx := context.Background()

	_, err := svc.AddProduct(ctx, "p1", "1", "Product", 10.0)
	assert.Error(t, err) // store doesn't exist yet

	_, _ = svc.CreateStore(ctx, "1", "Test Store")

	product, err := svc.AddProduct(ctx, "p1", "1", "Product", 10.0)
	assert.NoError(t, err)
	assert.Equal(t, "Product", product.Name)
	assert.Equal(t, 10.0, product.Price)
}
