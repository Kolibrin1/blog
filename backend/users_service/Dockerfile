FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o users-service ./cmd/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/users-service /usr/local/bin/users-service
CMD ["/usr/local/bin/users-service"]
