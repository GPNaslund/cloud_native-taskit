package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/dto"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for update task route
type UpdateTaskHandler struct {
	dataLayer     pb.DataLayerClient
	messageBroker interfaces.IMessageBroker
}

// Constructor function
func NewUpdateTaskHandler(dataLayer pb.DataLayerClient, messageBroker interfaces.IMessageBroker) *UpdateTaskHandler {
	return &UpdateTaskHandler{
		dataLayer:     dataLayer,
		messageBroker: messageBroker,
	}
}

// Function for handling incoming requests
func (u *UpdateTaskHandler) Handle(ctx *fiber.Ctx) error {
	// Get user id from locals set by middleware
	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		log.Printf("Update task handler got request without user id")
		return u.internalServerError(ctx)
	}

	// Extract task id from url parameter
	taskId := ctx.Params("taskId")

	if userId == "" || taskId == "" {
		log.Printf("Update task handler got request with userId: %s and taskId: %s", userId, taskId)
		return u.internalServerError(ctx)
	}

	// Parse request body
	task := new(dto.Task)
	if err := ctx.BodyParser(task); err != nil {
		return err
	}

	// Create grpc read task request
	var readOpts []grpc.CallOption
	readReq := pb.ReadTaskRequest{UserId: userId, TaskId: taskId}
	readRes, err := u.dataLayer.ReadTask(ctx.Context(), &readReq, readOpts...)

	if err != nil {
		log.Printf("Read task grpc call got error: %s", err.Error())
		return u.internalServerError(ctx)
	}

	// Get task pre update
	if readRes.Status != pb.DataStatus_Data_Success {
		if readRes.Status == pb.DataStatus_Data_No_Task_Found {
			return ctx.Status(fiber.StatusBadRequest).SendString("No task with provided id found")
		} else {
			return u.internalServerError(ctx)
		}
	}

	// Create grpc update task request
	var opts []grpc.CallOption
	taskDto := pb.TaskDTO{TaskId: taskId, Title: task.Title, Details: task.Details, IsDone: task.IsDone}
	req := pb.UpdateTaskRequest{UserId: userId, Task: &taskDto}

	// Call grpc update task function
	res, err := u.dataLayer.UpdateTask(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Update task grpc call got error: %s", err.Error())
		return u.internalServerError(ctx)
	}

	// Return the updated task on success
	if res.Status == pb.DataStatus_Data_Success {

		// Publish message to broker if status has changed
		if readRes.Task.IsDone != task.IsDone {
			if task.IsDone == true {
				err = u.messageBroker.PublishTaskCompleted(ctx.Context(), task.Title)
				if err != nil {
					log.Printf("Failed to publish task completed: %s", err.Error())
				}
			} else {
				err = u.messageBroker.PublishTaskUncompleted(ctx.Context(), task.Title)
				if err != nil {
					log.Printf("Failed to publish task uncompleted: %s", err.Error())
				}
			}
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":      res.Task.TaskId,
			"title":   res.Task.Title,
			"details": res.Task.Details,
			"is_done": res.Task.IsDone,
		})
	} else if res.Status == pb.DataStatus_Data_No_Task_Found {
		return ctx.Status(fiber.StatusBadRequest).SendString("No task with provided id found")
	} else {
		log.Printf("Update task grpc call got error: %s", res.Status.String())
		return u.internalServerError(ctx)
	}
}

// Helper function for returning internal server error
func (u *UpdateTaskHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
}
