package internal

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Represents a task message from the RabbitMQ channel
type TaskMessage struct {
	MessageType string `json:"message-type"`
	TaskTitle   string `json:"task-title"`
	Username    string `json:"username"`
}

// Interface representing the messenger to slack
type ISlackMessenger interface {
	PostMessage(message string) error
}

// Represents the Notification service
type NotificationService struct {
	slackMessenger ISlackMessenger
}

// Constructor function
func NewNotificationService(slackMessenger ISlackMessenger) *NotificationService {
	return &NotificationService{
		slackMessenger: slackMessenger,
	}
}

// Main function to start the notification service, and listen to messages on the queue + channel
func (n *NotificationService) Run(amqpConnString string, taskNotificationQueueName string) {
	// Create connection to RabbitMQ
	conn, err := amqp.Dial(amqpConnString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err.Error())
	}
	defer conn.Close()

	// Open channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel to RabbitMQ: %s", err.Error())
	}

	err = channel.Qos(
		1, // Handle one message at a time
		0,
		false,
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err.Error())
	}

	// Setup queue, same as publisher
	queue, err := channel.QueueDeclare(
		taskNotificationQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err.Error())
	}

	// Set up consume parameters
	messages, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume channel: %s", err.Error())
	}

	var forever chan struct{}

	go func() {
		for message := range messages {
			if err := n.handleIncomingMessage(message); err != nil {
				message.Nack(false, true)
				continue
			}
			message.Ack(false)
		}
	}()

	log.Printf("Ready for task update messages..")
	<-forever
}

// Helper function for handling incoming messages
func (n *NotificationService) handleIncomingMessage(message amqp.Delivery) error {
	var taskMessage TaskMessage
	err := json.Unmarshal(message.Body, &taskMessage)
	if err != nil {
		log.Printf("Failed to unmarshal message body: %s", err.Error())
		return err
	}

	// Validate contents
	if taskMessage.TaskTitle == "" || taskMessage.MessageType == "" || taskMessage.Username == "" {
		return fmt.Errorf("invalid message: missing required fields")
	}

	formattedMsg := fmt.Sprintf("%s was %s by %s", taskMessage.TaskTitle, taskMessage.MessageType, taskMessage.Username)
	if err := n.slackMessenger.PostMessage(formattedMsg); err != nil {
		return fmt.Errorf("failed to post to Slack: %w", err)
	}

	log.Printf("Successfully processed message: %s", formattedMsg)
	return nil
}
