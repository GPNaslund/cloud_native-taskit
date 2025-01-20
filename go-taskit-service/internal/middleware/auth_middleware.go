package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the authentication middleware
type AuthMiddleware struct {
	dataLayer     pb.DataLayerClient
	sessionLayer  pb.SessionLayerClient
	authService   interfaces.IAuthService
	cookieService interfaces.ICookieService
}

// Constructor function
func NewAuthMiddleware(dataLayer pb.DataLayerClient, sessionLayer pb.SessionLayerClient, authService interfaces.IAuthService, cookieService interfaces.ICookieService) *AuthMiddleware {
	return &AuthMiddleware{
		dataLayer:     dataLayer,
		sessionLayer:  sessionLayer,
		authService:   authService,
		cookieService: cookieService,
	}
}

func (a *AuthMiddleware) Authenticate(ctx *fiber.Ctx) error {
	userId := a.cookieService.GetUserIdFromCookie(ctx)
	if userId == "" {
		log.Println("No cookie provided")
		return ctx.Status(fiber.StatusUnauthorized).SendString("No cookie provided")
	}

	var validateSessionOpts []grpc.CallOption
	validateSessionReq := pb.GetSessionRequest{UserId: userId}

	validatedSession, err := a.sessionLayer.GetSession(ctx.Context(), &validateSessionReq, validateSessionOpts...)
	if err != nil {
		log.Printf("Get session gRPC call returned an error: %s", err.Error())
		return a.internalServerError(ctx)
	}

	switch validatedSession.Status {
	case pb.SessionStatus_Session_Success:
		if validatedSession.SessionToken != nil && *validatedSession.SessionToken == a.cookieService.GetSessionTokenFromCookie(ctx) {
			ctx.Locals("userId", userId)
			return ctx.Next()
		}
		log.Println("Session token mismatch")
		a.cookieService.SetDeletionCookie(ctx)
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid session")

	case pb.SessionStatus_Session_No_Session_Found:
		log.Println("No valid session found")
		a.cookieService.SetDeletionCookie(ctx)
		return ctx.Status(fiber.StatusUnauthorized).SendString("Session expired")
	default:
		return a.internalServerError(ctx)
	}
}

// Helper method for returning internal server error
func (a *AuthMiddleware) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Something went wrong internally, try again later!")
}
