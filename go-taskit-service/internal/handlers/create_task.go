package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/dto"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Handler for create task route
type CreateTaskHandler struct {
	dataLayer     pb.DataLayerClient
	messageBroker interfaces.IMessageBroker
}

// Constructor function
func NewCreateTaskHandler(dataLayer pb.DataLayerClient, messageBroker interfaces.IMessageBroker) *CreateTaskHandler {
	return &CreateTaskHandler{
		dataLayer:     dataLayer,
		messageBroker: messageBroker,
	}
}

// Method for handling incoming request
func (c *CreateTaskHandler) Handle(ctx *fiber.Ctx) error {
	log.Println("Handling create task request")
	// Get user id from request locals
	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		log.Printf("Create task handler got request without userId")
		return c.internalServerError(ctx)
	}

	// Parse incoming request body
	task := new(dto.Task)
	if err := ctx.BodyParser(&task); err != nil {
		return err
	}

	// Create grpc create task request
	var opts []grpc.CallOption
	taskDto := pb.TaskDTO{TaskId: task.Id, Title: task.Title, Details: task.Details, IsDone: task.IsDone}
	req := pb.CreateTaskRequest{UserId: userId, Task: &taskDto}

	// Call data layer to create a new task
	res, err := c.dataLayer.CreateTask(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Create task grpc service returned error: %s", err.Error())
		return c.internalServerError(ctx)
	}

	// Return the created task on success
	if res.Status == pb.DataStatus_Data_Success {
		err = c.messageBroker.PublishTaskCreated(ctx.Context(), res.Task.Title)
		if err != nil {
			log.Printf("Failed to publish task completed: %s", err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":      res.Task.TaskId,
			"title":   res.Task.Title,
			"details": res.Task.Details,
			"is_done": res.Task.IsDone,
		})
	} else {
		log.Printf("Create task grpc service returned status: %s", res.Status.String())
		return c.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (c *CreateTaskHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
}
