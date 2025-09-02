package repo

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
	"stores-lite/internal/domain"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(url string) (*RedisClient, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return &RedisClient{rdb: redis.NewClient(opt)}, nil
}

func (c *RedisClient) Close() error { return c.rdb.Close() }

func (c *RedisClient) CacheProduct(ctx context.Context, p *domain.Product) error {
	b, _ := json.Marshal(p)
	return c.rdb.Set(ctx, c.key(p.ID), b, 0).Err()
}

func (c *RedisClient) GetCachedProduct(ctx context.Context, id int64) (*domain.Product, error) {
	s, err := c.rdb.Get(ctx, c.key(id)).Result()
	if err != nil {
		return nil, err
	}
	var p domain.Product
	_ = json.Unmarshal([]byte(s), &p)
	return &p, nil
}

func (c *RedisClient) PublishOrder(ctx context.Context, o *domain.Order) error {
	b, _ := json.Marshal(o)
	return c.rdb.Publish(ctx, "orders.created", b).Err()
}

func (c *RedisClient) key(id int64) string {
	return strings.Join([]string{"product", "v1", string(rune(id))}, ":")
}
