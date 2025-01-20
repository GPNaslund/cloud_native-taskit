package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gn222gq.2dv013.a2/internal"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Entry point for session service
func main() {
	redisURL := os.Getenv("REDIS_CONNECTION_STRING")
	if redisURL == "" {
		log.Fatal("REDIS_CONNECTION_STRING must be set")
	}

	address := os.Getenv("TASKIT_SESSIONSERVICE_ADDRESS")
	if address == "" {
		log.Fatal("TASKIT_SESSIONSERVICE_ADDRESS must be set")
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("")
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse redis url: %s", err.Error())
	}

	client := redis.NewClient(opt)
	ttl := 1 * time.Hour
	sessionLayerServer := internal.NewSessionLayerServer(client, ttl)
	pb.RegisterSessionLayerServer(grpcServer, sessionLayerServer)
	grpcServer.Serve(lis)
}
