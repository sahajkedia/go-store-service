package store

import (
	"context"
	"errors"
)

type Store struct {
	ID   string
	Name string
}

type Product struct {
	ID      string
	StoreID string
	Name    string
	Price   float64
}

type Service struct {
	stores   map[string]*Store
	products map[string]*Product
}

func NewService() *Service {
	return &Service{
		stores:   make(map[string]*Store),
		products: make(map[string]*Product),
	}
}

func (s *Service) CreateStore(ctx context.Context, id string, name string) (*Store, error) {
	if _, exists := s.stores[id]; exists {
		return nil, errors.New("store already exists")
	}
	store := &Store{ID: id, Name: name}
	s.stores[id] = store
	return store, nil
}

func (s *Service) AddProduct(ctx context.Context, id string, storeID string, name string, price float64) (*Product, error) {
	if _, ok := s.stores[storeID]; !ok {
		return nil, errors.New("store does not exist")
	}
	product := &Product{ID: id, StoreID: storeID, Name: name, Price: price}
	s.products[id] = product
	return product, nil
}
