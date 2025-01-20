package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler of the refresh session route
type RefreshSessionHandler struct {
	cookieService interfaces.ICookieService
	sessionClient pb.SessionLayerClient
}

// Constructor function
func NewRefreshSessionHandler(cookieService interfaces.ICookieService, sessionClient pb.SessionLayerClient) *RefreshSessionHandler {
	return &RefreshSessionHandler{
		cookieService: cookieService,
		sessionClient: sessionClient,
	}
}

// Function for handling incoming requests
func (r *RefreshSessionHandler) Handle(ctx *fiber.Ctx) error {
	log.Println("Handling refresh session request..")

	// Extract user id from provided cookie
	userId := r.cookieService.GetUserIdFromCookie(ctx)
	if userId == "" {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No cookie with user id provided")
	}

	// Check for valid existing session
	var opts []grpc.CallOption
	req := pb.GetSessionRequest{UserId: userId}
	res, err := r.sessionClient.GetSession(ctx.Context(), &req, opts...)
	if err != nil {
		log.Printf("Get session grpc call returned error: %s", err.Error())
		return r.internalServerError(ctx)
	}

	// If valid session exists, create a new session and return a new cookie
	if res.Status == pb.SessionStatus_Session_Success {
		createReq := pb.CreateSessionRequest{UserId: userId}
		createRes, err := r.sessionClient.CreateSession(ctx.Context(), &createReq, opts...)
		if err != nil {
			log.Printf("Create session gRPC call returned error: %s", err.Error())
			return r.internalServerError(ctx)
		}
		if createRes.Status == pb.SessionStatus_Session_Success {
			r.cookieService.SetSessionCookie(ctx, *createRes.SessionToken, userId)
			return ctx.Status(fiber.StatusOK).SendString("Session refreshed")
		} else {
			log.Printf("Create session gRPC call returned status: %s", createRes.Status.String())
			return r.internalServerError(ctx)
		}
	} else if res.Status == pb.SessionStatus_Session_No_Session_Found {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No valid session found")
	} else {
		log.Printf("Get session gRPC call returned status: %s", res.Status.String())
		return r.internalServerError(ctx)
	}
}

// Helper method for returning internal server error response
func (r *RefreshSessionHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
}
