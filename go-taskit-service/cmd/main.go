package main

import (
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gn222gq.2dv013.a2/internal"
	"gn222gq.2dv013.a2/internal/handlers"
	service "gn222gq.2dv013.a2/internal/services"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Entry point of application
func main() {
	// Reads data service adress environment variable
	dataServiceAddress := os.Getenv("TASKIT_DATASERVICE_ADDRESS")
	if dataServiceAddress == "" {
		log.Fatal("TASKIT_DATASERVICE_ADDRESS must be set")
	}

	// Reads session service address environment variable
	sessionServiceAddress := os.Getenv("TASKIT_SESSIONSERVICE_ADDRESS")
	if sessionServiceAddress == "" {
		log.Fatal("TASKIT_SESSIONSERVICE_ADDRESS must be set")
	}

	// Get the listening port for the application
	listeningPort := os.Getenv("TASKIT_SERVICE_PORT")
	if listeningPort == "" {
		log.Fatal("TASKIT_SERVICE_PORT must be set")
	}

	// Get the base url for the application
	baseUrl := os.Getenv("HOST_BASE_URL")
	if baseUrl == "" {
		log.Fatal("HOST_BASE_URL must be set")
	}

	amqpConnString := os.Getenv("AMQP_CONN_STRING")
	if amqpConnString == "" {
		log.Fatal("AMQP_CONN_STRING must be set")
	}

	taskNotificationQueueName := os.Getenv("TASK_NOTIFICATION_QUEUE_NAME")
	if taskNotificationQueueName == "" {
		log.Fatal("TASK_NOTIFICATION_QUEUE_NAME must be set")
	}

	// Creates a new grpc client connection for data service
	dataConn, err := grpc.NewClient(dataServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error creating grpc client for data service: %s", err.Error())
	}

	// Creates a new grpc client connection for session service
	sessionConn, err := grpc.NewClient(sessionServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error creating grpc client for session service: %s", err.Error())
	}

	// Instansiate services and grpc clients
	dataLayer := pb.NewDataLayerClient(dataConn)
	sessionLayer := pb.NewSessionLayerClient(sessionConn)
	cookieService := service.NewCookieService("taskit_session", 1*time.Hour)
	authService := service.NewAuthService()

	conn, err := amqp.Dial(amqpConnString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	messageBroker := service.NewMessageBroker(conn, taskNotificationQueueName)

	// Creates handler factory
	factory := handlers.NewHandlerFactory(dataLayer, sessionLayer, cookieService, authService, messageBroker)

	// Creates the router
	router := internal.NewRouter(listeningPort, baseUrl, factory)

	// Starts the router
	router.Start()
}
