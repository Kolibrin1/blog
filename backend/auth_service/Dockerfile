FROM golang:1.23.2-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o auth-service ./cmd/main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/auth-service /usr/local/bin/auth-service
CMD ["/usr/local/bin/auth-service"]
