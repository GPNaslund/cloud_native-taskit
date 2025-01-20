package internal

import (
	"log"

	"github.com/slack-go/slack"
)

// Represents a slack messenger
type SlackMessenger struct {
	client    *slack.Client
	channelId string
}

// Constructor method
func NewSlackMessenger(slackToken string, channelId string) *SlackMessenger {
	client := slack.New(slackToken)
	return &SlackMessenger{
		client:    client,
		channelId: channelId,
	}
}

// Function for posting messages to slack channel
func (s *SlackMessenger) PostMessage(message string) error {
	channelId, timestamp, err := s.client.PostMessage(
		s.channelId,
		slack.MsgOptionText(message, false))
	if err != nil {
		log.Printf("Failed to send message to slack channel: %s", err.Error())
		return err
	}

	log.Printf("Message sent successfully to channel %s at %s", channelId, timestamp)
	return nil
}
