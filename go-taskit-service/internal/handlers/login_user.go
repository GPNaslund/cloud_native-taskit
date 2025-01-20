package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/dto"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for the login route
type LoginHandler struct {
	dataLayer     pb.DataLayerClient
	sessionLayer  pb.SessionLayerClient
	authService   interfaces.IAuthService
	cookieService interfaces.ICookieService
}

// Constructor function
func NewLoginHandler(dataLayer pb.DataLayerClient, sessionLayer pb.SessionLayerClient, authService interfaces.IAuthService, cookieService interfaces.ICookieService) *LoginHandler {
	return &LoginHandler{
		dataLayer:     dataLayer,
		sessionLayer:  sessionLayer,
		authService:   authService,
		cookieService: cookieService,
	}
}

// Function for handling incoming requests
func (l *LoginHandler) Handle(ctx *fiber.Ctx) error {
	// Extract request body
	var user dto.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	// Create read user grpc request
	var getUserOpts []grpc.CallOption
	getUserReq := pb.ReadUserRequest{Username: user.Username}

	// Make grpc read user request to data layer
	getUserRes, err := l.dataLayer.ReadUser(ctx.Context(), &getUserReq, getUserOpts...)
	if err != nil {
		log.Printf("Get user grpc call returned an error: %s", err.Error())
		return l.internalServerError(ctx)
	}

	// If user exists, compare passwords
	if getUserRes.Status == pb.DataStatus_Data_Success {
		err = l.authService.ComparePasswords(user.Password, getUserRes.User.Password)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
		}

		// Create a grpc create session request
		var opts []grpc.CallOption
		req := pb.CreateSessionRequest{UserId: *getUserRes.User.UserId}

		// Call grpc create session function
		res, err := l.sessionLayer.CreateSession(ctx.Context(), &req, opts...)
		if err != nil {
			log.Printf("Create session grpc call returned error: %s", err.Error())
			return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later")
		}

		// If session creation is successful return cookie with token
		if res.Status == pb.SessionStatus_Session_Success {
			l.cookieService.SetSessionCookie(ctx, *res.SessionToken, *getUserRes.User.UserId)
			return ctx.Status(fiber.StatusOK).SendString("Login successful")
		} else {
			log.Printf("Create session grpc call returned status: %s", res.Status.String())
			return l.internalServerError(ctx)
		}

	} else if getUserRes.Status == pb.DataStatus_Data_No_User_Found {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No user found with provided username")
	} else {
		log.Printf("Get user grpc call returned status code: %s", getUserRes.Status.String())
		return l.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (l *LoginHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later!")
}
