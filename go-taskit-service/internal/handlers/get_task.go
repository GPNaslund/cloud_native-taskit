package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Handler for get task route
type GetTaskHandler struct {
	dataLayer pb.DataLayerClient
}

// Constructor function
func NewGetTaskHandler(dataLayer pb.DataLayerClient) *GetTaskHandler {
	return &GetTaskHandler{
		dataLayer: dataLayer,
	}
}

// Function for handling incoming requests
func (g *GetTaskHandler) Handle(ctx *fiber.Ctx) error {
	// Extract user id from request locals
	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		log.Printf("Get task handler got request without user id")
		return g.internalServerError(ctx)
	}

	// Extract task id from request parameters
	taskId := ctx.Params("taskId")
	if taskId == "" {
		log.Printf("Get task handler got request with taskId: %s", taskId)
		return g.internalServerError(ctx)
	}

	// Creates read task request
	var opts []grpc.CallOption
	req := pb.ReadTaskRequest{UserId: userId, TaskId: taskId}

	// Call datalayer read task function
	res, err := g.dataLayer.ReadTask(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Read task grpc call returned error: %s", err.Error())
		return g.internalServerError(ctx)
	}

	// Return task on success
	if res.Status == pb.DataStatus_Data_Success {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":      res.Task.TaskId,
			"title":   res.Task.Title,
			"details": res.Task.Details,
			"is_done": res.Task.IsDone,
		})
	} else if res.Status == pb.DataStatus_Data_No_Task_Found {
		return ctx.Status(fiber.StatusBadRequest).SendString("Task with provided id not found")
	} else {
		log.Printf("Read task grpc call returned status: %s", res.Status.String())
		return g.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (g *GetTaskHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
}
