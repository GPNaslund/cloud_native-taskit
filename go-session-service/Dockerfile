# Build
FROM golang:1.23.3-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org
RUN echo "hosts: files dns" > /etc/nsswitch.conf
RUN go clean -modcache && go mod download
COPY . .
WORKDIR /app/cmd
RUN go build -o session-service .

# Runtime
FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/cmd/session-service .
EXPOSE 50052

# Run the application
ENTRYPOINT ["./session-service"]
