package service

import (
	"context"
	"testing"

	"stores-lite/internal/domain"
)

type stubRepo struct {
	products []domain.Product
}

func (s *stubRepo) CreateProduct(ctx context.Context, p *domain.Product) error {
	p.ID = int64(len(s.products) + 1)
	s.products = append(s.products, *p)
	return nil
}
func (s *stubRepo) ListProducts(ctx context.Context) ([]domain.Product, error) { return s.products, nil }
func (s *stubRepo) GetProduct(ctx context.Context, id int64) (*domain.Product, error) {
	for _, p := range s.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}
type stubOrderRepo struct{}
func (s *stubOrderRepo) CreateOrder(ctx context.Context, o *domain.Order) error { o.ID = 1; return nil }

type stubCache struct{}
func (s *stubCache) CacheProduct(ctx context.Context, p *domain.Product) error { return nil }
func (s *stubCache) PublishOrder(ctx context.Context, o *domain.Order) error   { return nil }

func TestCreateProduct(t *testing.T) {
	pg := &stubRepo{}
	cache := &stubCache{}
	s := &Service{products: pg, orders: &stubOrderRepo{}, cache: nil}
	ctx := context.Background()
	p, err := s.CreateProduct(ctx, "Sticker", 999, 100)
	if err != nil { t.Fatal(err) }
	if p.ID != 1 || p.PriceCents != 999 { t.Fatal("unexpected product data") }
}

func TestCheckout(t *testing.T) {
	pg := &stubRepo{products: []domain.Product{{ID: 1, Name: "Sticker", PriceCents: 500, Stock: 10}}}
	s := &Service{products: pg, orders: &stubOrderRepo{}, cache: nil}
	ctx := context.Background()
	o, err := s.Checkout(ctx, 1, 2)
	if err != nil { t.Fatal(err) }
	if o.Amount != 1000 { t.Fatal("wrong amount") }
}
