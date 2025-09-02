package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"stores-lite/internal/domain"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, url string) (*PostgresRepo, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{db: pool}, nil
}

func (r *PostgresRepo) Close() { r.db.Close() }

func (r *PostgresRepo) CreateProduct(ctx context.Context, p *domain.Product) error {
	q := `INSERT INTO products(name, price_cents, stock) VALUES($1,$2,$3) RETURNING id`
	return r.db.QueryRow(ctx, q, p.Name, p.PriceCents, p.Stock).Scan(&p.ID)
}

func (r *PostgresRepo) ListProducts(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, price_cents, stock FROM products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.PriceCents, &p.Stock); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (r *PostgresRepo) GetProduct(ctx context.Context, id int64) (*domain.Product, error) {
	var p domain.Product
	err := r.db.QueryRow(ctx, `SELECT id, name, price_cents, stock FROM products WHERE id=$1`, id).Scan(
		&p.ID, &p.Name, &p.PriceCents, &p.Stock,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PostgresRepo) CreateOrder(ctx context.Context, o *domain.Order) error {
	q := `INSERT INTO orders(product_id, qty, amount) VALUES($1,$2,$3) RETURNING id`
	return r.db.QueryRow(ctx, q, o.ProductID, o.Qty, o.Amount).Scan(&o.ID)
}

// Interfaces for service
type ProductRepo interface {
	CreateProduct(context.Context, *domain.Product) error
	ListProducts(context.Context) ([]domain.Product, error)
	GetProduct(context.Context, int64) (*domain.Product, error)
}

type OrderRepo interface {
	CreateOrder(context.Context, *domain.Order) error
}

var _ ProductRepo = (*PostgresRepo)(nil)
var _ OrderRepo = (*PostgresRepo)(nil)
