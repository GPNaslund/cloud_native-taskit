package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/dto"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for getting multiple tasks route
type GetMultipleTasksHandler struct {
	dataLayer pb.DataLayerClient
}

// Constuctor function
func NewGetMultipleTasksHandler(dataLayer pb.DataLayerClient) *GetMultipleTasksHandler {
	return &GetMultipleTasksHandler{
		dataLayer: dataLayer,
	}
}

// Function for handling incoming requests
func (g *GetMultipleTasksHandler) Handle(ctx *fiber.Ctx) error {
	// Extract the user id from request locals
	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		log.Printf("Get multiple tasks handler got request without user id")
		return g.internalServerError(ctx)
	}

	// TODO: CHECK IF GRPC CALL WILL RETURNED PAGINATED RESULTS!
	limit := ctx.Query("limit", "25")
	page := ctx.Query("page", "1")

	// Check if limit is number
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Limit query parameter must be a number")
	}

	// Check if page is number
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Page query parameter must be a number")
	}

	// Create grpc multiple tasks request
	var opts []grpc.CallOption
	req := pb.ReadMultipleTasksRequest{UserId: userId, Limit: int32(limitInt), Page: int32(pageInt)}

	// Send the read multiple tasks request
	res, err := g.dataLayer.ReadMultipleTasks(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Read multiple tasks grpc call returned error: %s", err.Error())
		return g.internalServerError(ctx)
	}

	// Return tasks on success
	if res.Status == pb.DataStatus_Data_Success {
		taskDtoSlice := res.Tasks
		taskResults := []dto.Task{}
		for _, task := range taskDtoSlice {
			taskResults = append(taskResults, dto.Task{Id: task.TaskId, Title: task.Title, Details: task.Details, IsDone: task.IsDone})
		}
		return ctx.Status(fiber.StatusOK).JSON(taskResults)
	} else {
		log.Printf("Read multiple tasks grpc call returned status: %s", res.Status.String())
		return g.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (g *GetMultipleTasksHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later!")
}
