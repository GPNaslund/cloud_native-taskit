package main

import (
	"context"
	"log"
	"net"
	"os"

	"gn222gq.2dv013.a2/internal"
	"gn222gq.2dv013.a2/internal/dto"
	service "gn222gq.2dv013.a2/internal/services"
	pb "gn222gq.2dv013.a2/protos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// Entry point for the grpc data service server
func main() {
	address := os.Getenv("TASKIT_DATASERVICE_ADDRESS")
	if address == "" {
		log.Fatalf("TASKIT_DATASERVICE_ADDRESS must be set")
	}

	mongoURI := os.Getenv("TASKIT_MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("TASKIT_MONGODB_URI must be set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	dbDetails := dto.DatabaseDetails{DatabaseName: "taskit", UserCollectionName: "users"}
	userService := service.NewUserService(client, dbDetails)
	taskService := service.NewTaskService(client, dbDetails)
	dataLayerService := internal.NewDataLayerServer(userService, taskService)

	pb.RegisterDataLayerServer(grpcServer, dataLayerService)
	grpcServer.Serve(lis)
}
