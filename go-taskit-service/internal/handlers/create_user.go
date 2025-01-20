package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/dto"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Interface with functionality for hashing a password
type PasswordHasher interface {
	HashPassword(password string) ([]byte, error)
}

// Represents the handler for create user route
type CreateUserHandler struct {
	hasher    PasswordHasher
	dataLayer pb.DataLayerClient
}

// Constructor function
func NewCreateUserHandler(passwordHasher PasswordHasher, dataLayerClient pb.DataLayerClient) *CreateUserHandler {
	return &CreateUserHandler{
		hasher:    passwordHasher,
		dataLayer: dataLayerClient,
	}
}

// Function for handling incoming request
func (c *CreateUserHandler) Handle(ctx *fiber.Ctx) error {
	log.Println("Handling create user request")
	// Parse request body for new user
	userData := new(dto.User)
	if err := ctx.BodyParser(&userData); err != nil {
		return err
	}

	// Hashes password
	hashed, err := c.hasher.HashPassword(userData.Password)
	if err != nil {
		log.Printf("Error occured on password hashing: %s", err.Error())
		return c.internalServerError(ctx)
	}

	// Creates the create user request
	req := pb.CreateUserRequest{Username: userData.Username, Password: string(hashed)}
	var opts []grpc.CallOption

	// Makes data layer create user request
	result, err := c.dataLayer.CreateUser(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Error occured when calling data grpc: %s", err.Error())
		return c.internalServerError(ctx)
	}

	// Return the user on successful creation
	if result.Status == pb.DataStatus_Data_Success {
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"username": result.User.Username,
			"user_id":  result.User.UserId,
		})
	} else if result.Status == pb.DataStatus_Data_Invalid_Username {
		return ctx.Status(fiber.StatusBadRequest).SendString("Username is not valid, try another one")
	} else {
		log.Printf("Create user grpc call returned error: %s", result.Status.String())
		return c.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (c *CreateUserHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later!")
}
