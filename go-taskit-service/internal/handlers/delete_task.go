package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for delete task route
type DeleteTaskHandler struct {
	dataLayer pb.DataLayerClient
}

// Constructor function
func NewDeleteTaskHandler(dataLayer pb.DataLayerClient) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		dataLayer: dataLayer,
	}
}

// Function for handling incoming request
func (d *DeleteTaskHandler) Handle(ctx *fiber.Ctx) error {
	log.Printf("Handling delete request")

	// Extract user id from request locals set by auth middleware
	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		log.Printf("Delete task handler got request without user id")
		return d.internalServerError(ctx)
	}

	// Extract task id from url parameter
	taskId := ctx.Params("taskId")
	if taskId == "" {
		log.Printf("Delete task handler got empty taskId")
		return ctx.Status(fiber.StatusBadRequest).SendString("Task id cannot be empty")
	}

	// Create delete task request
	var opts []grpc.CallOption
	req := pb.DeleteTaskRequest{UserId: userId, TaskId: taskId}

	// Call data layer grpc delete task
	res, err := d.dataLayer.DeleteTask(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Delete task gRPC call returned error: %s", err.Error())
		return d.internalServerError(ctx)
	}

	switch res.Status {
	case pb.DataStatus_Data_Success:
		return ctx.Status(fiber.StatusNoContent).SendString("Deleted task successfully")
	case pb.DataStatus_Data_No_Task_Found:
		return ctx.Status(fiber.StatusNotFound).SendString("No task found")
	case pb.DataStatus_Data_No_User_Found:
		return ctx.Status(fiber.StatusNotFound).SendString("The owner of the task not found")
	default:
		log.Printf("Delete task gRPC call returned unexpected status: %s", res.Status)
		return d.internalServerError(ctx)
	}
}

// Helper method for returning internal server error.
func (d *DeleteTaskHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
}
