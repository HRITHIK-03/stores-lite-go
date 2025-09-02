# Build
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Run
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/api /app/api
EXPOSE 8080
ENTRYPOINT ["/app/api"]
