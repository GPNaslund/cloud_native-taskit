package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for validating a session
type ValidateSessionHandler struct {
	sessionLayer  pb.SessionLayerClient
	cookieService interfaces.ICookieService
}

// Constructor function
func NewValidateSessionHandler(sessionLayer pb.SessionLayerClient, cookieService interfaces.ICookieService) *ValidateSessionHandler {
	return &ValidateSessionHandler{
		sessionLayer:  sessionLayer,
		cookieService: cookieService,
	}
}

// Function for handling incoming requests
func (v *ValidateSessionHandler) Handle(ctx *fiber.Ctx) error {
	// Extract user id from the cookie
	userId := v.cookieService.GetUserIdFromCookie(ctx)
	if userId == "" {
		log.Printf("No session cookie found")
		return ctx.Status(fiber.StatusUnauthorized).SendString("No valid cookie provided")
	}

	// Create the grpc get session request
	var opts []grpc.CallOption
	req := pb.GetSessionRequest{UserId: userId}

	// Call the grpc get session function
	res, err := v.sessionLayer.GetSession(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Get session gRPC call returned error: %s", err.Error())
		return v.internalServerError(ctx)
	}

	switch res.Status {
	case pb.SessionStatus_Session_Success:
		// RENEW SESSION HERE!
		return ctx.SendStatus(fiber.StatusOK)
	case pb.SessionStatus_Session_No_Session_Found:
		log.Printf("No valid session found for user: %s", userId)
		return ctx.Status(fiber.StatusUnauthorized).SendString("No valid session found")
	default:
		log.Printf("Get session returned unexpected status: %v", res.Status)
		return v.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (v *ValidateSessionHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong internally, try again later",
	})
}
