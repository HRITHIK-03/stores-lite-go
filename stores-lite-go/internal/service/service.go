package service

import (
	"context"
	"errors"

	"stores-lite/internal/domain"
	"stores-lite/internal/repo"
)

type Service struct {
	products repo.ProductRepo
	orders   repo.OrderRepo
	cache    *repo.RedisClient
}

func New(pg *repo.PostgresRepo, cache *repo.RedisClient) *Service {
	return &Service{products: pg, orders: pg, cache: cache}
}

func (s *Service) CreateProduct(ctx context.Context, name string, price int64, stock int64) (*domain.Product, error) {
	if name == "" || price <= 0 {
		return nil, errors.New("invalid input")
	}
	p := &domain.Product{Name: name, PriceCents: price, Stock: stock}
	if err := s.products.CreateProduct(ctx, p); err != nil {
		return nil, err
	}
	_ = s.cache.CacheProduct(ctx, p)
	return p, nil
}

func (s *Service) ListProducts(ctx context.Context) ([]domain.Product, error) {
	return s.products.ListProducts(ctx)
}

func (s *Service) Checkout(ctx context.Context, productID, qty int64) (*domain.Order, error) {
	p, err := s.products.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	if qty <= 0 || p.Stock < qty {
		return nil, errors.New("insufficient stock")
	}
	o := &domain.Order{ProductID: productID, Qty: qty, Amount: p.PriceCents * qty}
	if err := s.orders.CreateOrder(ctx, o); err != nil {
		return nil, err
	}
	_ = s.cache.PublishOrder(ctx, o)
	return o, nil
}
