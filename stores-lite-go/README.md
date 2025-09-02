# stores-lite (Go + GraphQL + Postgres + Redis)
A compact, production-style sample that mirrors Sticker Mule's stack: **Go**, **TypeScript client (optional)**, **GraphQL**, **Postgres**, **Redis**, **Docker**, and **Kubernetes** manifests.

## What it showcases
- Clean architecture (handlers → services → repositories)
- **GraphQL** read API + **REST** write API
- **Postgres** with SQL migrations
- **Redis** for caching product lookups and publishing order events
- Request-scoped logging, graceful shutdown, env-based config
- Unit tests (service layer)
- **Dockerfile**, **docker-compose**, and minimal **K8s** manifests

## Quick start (Docker Compose)
```bash
docker compose up --build
# API: http://localhost:8080
# GraphQL: POST /graphql ; GraphiQL (dev): http://localhost:8080/graphiql
```

### Example GraphQL query
```graphql
query {
  products {
    id
    name
    priceCents
    stock
  }
}
```

### Example REST calls
```bash
# Create product
curl -X POST http://localhost:8080/api/products -H "Content-Type: application/json" -d '{"name":"Sticker Pack","priceCents":1299,"stock":100}'

# Checkout (creates order + publishes Redis event)
curl -X POST http://localhost:8080/api/checkout -H "Content-Type: application/json" -d '{"productId":1,"qty":2}'
```

## Stack
- Go 1.22, chi router, pgx, redis/v9, graphql-go
- Postgres 15
- Redis 7
- Docker + docker-compose
- K8s manifests (dev)

## Repo structure
```
.
├─ cmd/api/main.go
├─ internal/
│  ├─ config/config.go
│  ├─ db/migrations/001_init.sql
│  ├─ domain/models.go
│  ├─ repo/postgres.go
│  ├─ service/service.go
│  ├─ transport/
│  │  ├─ rest.go
│  │  ├─ graphql.go
│  │  └─ middleware.go
├─ go.mod
├─ docker-compose.yml
├─ Dockerfile
├─ k8s/deployment.yaml
└─ k8s/service.yaml
```
