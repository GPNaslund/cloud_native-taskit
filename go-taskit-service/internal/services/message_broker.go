package service

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskMessage struct {
	MessageType string `json:"message-type"`
	TaskTitle   string `json:"task-title"`
	Username    string `json:"username"`
}

type MessageBroker struct {
	conn      *amqp.Connection
	queueName string
}

func NewMessageBroker(conn *amqp.Connection, queueName string) *MessageBroker {
	return &MessageBroker{
		conn:      conn,
		queueName: queueName,
	}
}

func (m *MessageBroker) PublishTaskCompleted(ctx context.Context, taskTitle string) error {
	return m.publishMessage(ctx, TaskMessage{
		MessageType: "completed",
		TaskTitle:   taskTitle,
		Username:    "gn222gq",
	})
}

func (m *MessageBroker) PublishTaskCreated(ctx context.Context, taskTitle string) error {
	return m.publishMessage(ctx, TaskMessage{
		MessageType: "created",
		TaskTitle:   taskTitle,
		Username:    "gn222gq",
	})
}

func (m *MessageBroker) PublishTaskUncompleted(ctx context.Context, taskTitle string) error {
	return m.publishMessage(ctx, TaskMessage{
		MessageType: "uncompleted",
		TaskTitle:   taskTitle,
		Username:    "gn222gq",
	})
}

func (m *MessageBroker) publishMessage(ctx context.Context, taskMessage TaskMessage) error {
	channel, err := m.conn.Channel()
	if err != nil {
		log.Printf("Error while creating RabbitMQ channel: %s", err.Error())
		return err
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		m.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error while creating RabbitMQ queue: %s", err.Error())
		return err
	}

	buffer, err := json.Marshal(&taskMessage)
	if err != nil {
		log.Printf("Error while marshalling task message: %s", err.Error())
		return err
	}

	err = channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         buffer,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		log.Printf("Error while publishing message: %s", err.Error())
		return err
	}

	log.Printf("Successfully sent message: %s", buffer)
	return nil
}
