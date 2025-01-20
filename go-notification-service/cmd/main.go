package main

import (
	"log"
	"os"

	"gn222gq.2dv013.a2/internal"
)

// Entry point for service
func main() {
	// Get RabbitMQ connection string from env
	amqpConnString := os.Getenv("AMQP_CONN_STRING")
	if amqpConnString == "" {
		log.Fatal("AMQP_CONN_STRING must be set")
	}

	// Get task notification queue name from env
	taskNotificationQueueName := os.Getenv("TASK_NOTIFICATION_QUEUE_NAME")
	if taskNotificationQueueName == "" {
		log.Fatal("TASK_NOTIFICATION_QUEUE_NAME must be set")
	}

	// Get slack token from env
	slackToken := os.Getenv("SLACK_TOKEN")
	if slackToken == "" {
		log.Fatal("SLACK_TOKEN must be set")
	}

	// Get slack channel id from env
	slackChannelId := os.Getenv("SLACK_CHANNEL_ID")
	if slackChannelId == "" {
		log.Fatal("SLACK_CHANNEL_ID must be set")
	}

	// Create slack messenger service
	slackMessenger := internal.NewSlackMessenger(slackToken, slackChannelId)
	// Create notification service
	notficationService := internal.NewNotificationService(slackMessenger)

	// Start notification service
	notficationService.Run(amqpConnString, taskNotificationQueueName)

}
