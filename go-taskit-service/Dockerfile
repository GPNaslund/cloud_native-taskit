# Build
FROM golang:1.23.3-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org
RUN echo "hosts: files dns" > /etc/nsswitch.conf
RUN go clean -modcache && go mod download
COPY . .
WORKDIR /app/cmd
RUN go build -o taskit-service .

# Runtime
FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/cmd/taskit-service .
EXPOSE 8080

# Run the application
ENTRYPOINT ["./taskit-service"]
