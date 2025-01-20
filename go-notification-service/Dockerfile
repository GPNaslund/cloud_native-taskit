# Build
FROM golang:1.23.3-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN go build -o notification-service .

# Runtime
FROM debian:bullseye
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY --from=builder /app/cmd/notification-service .

# Run the application
ENTRYPOINT ["./notification-service"]