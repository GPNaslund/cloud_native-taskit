package interfaces

import "context"

type IMessageBroker interface {
	PublishTaskCompleted(ctx context.Context, taskTitle string) error
	PublishTaskCreated(ctx context.Context, taskTitle string) error
	PublishTaskUncompleted(ctx context.Context, taskTitle string) error
}
