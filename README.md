Stores Lite (Go + GraphQL + Postgres + Redis)










A production-style sample project showcasing a minimal commerce platform built with Go, GraphQL, Postgres, and Redis.
It demonstrates clean architecture, containerized deployment, and cloud-ready practices.

✨ Features

GraphQL API for product queries

REST API for commands (create product, checkout)

Postgres with SQL migrations for persistence

Redis for caching & publishing order events

Layered architecture: Handlers → Services → Repositories

Request-scoped logging, graceful shutdown, and env-based config

Dockerfile, docker-compose, and Kubernetes manifests

Unit tests for the service layer

🗂 Project Structure
.
├── cmd/api/main.go          # Application entrypoint
├── internal/
│   ├── config/              # Environment configuration
│   ├── db/migrations/       # Database migrations
│   ├── domain/              # Core models
│   ├── repo/                # Postgres & Redis implementations
│   ├── service/             # Business logic
│   └── transport/           # REST & GraphQL handlers, middleware
├── Dockerfile
├── docker-compose.yml
├── k8s/                     # Kubernetes manifests
└── README.md

🚀 Getting Started
Run with Docker Compose
docker compose up --build


REST API: http://localhost:8080/api

GraphQL endpoint: http://localhost:8080/graphql

GraphiQL (dev only): http://localhost:8080/graphiql

🔎 Example Usage

GraphQL Query

query {
  products {
    id
    name
    priceCents
    stock
  }
}


REST Requests

# Create a product
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Sticker Pack","priceCents":1299,"stock":100}'

# Checkout an order
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{"productId":1,"qty":2}'

🛠 Tech Stack

Go 1.22

chi
 router

pgx
 (Postgres driver)

go-redis v9

graphql-go

Docker
 & docker-compose

Kubernetes
 (deployment & service manifests)

✅ Running Tests
go test ./internal/service -v

🌍 Deployment

Dockerfile: multi-stage build with distroless runtime

docker-compose: local development with Postgres + Redis

Kubernetes manifests: Deployment + Service for cluster deployment

📌 About

This project is intended as a showcase repository for demonstrating backend development in Go with a modern stack (GraphQL, Postgres, Redis, Docker, Kubernetes).
It mirrors real commerce workflows (products, checkout) and highlights scalable, cloud-ready design.
