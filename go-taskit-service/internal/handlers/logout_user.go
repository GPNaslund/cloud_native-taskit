package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gn222gq.2dv013.a2/internal/interfaces"
	pb "gn222gq.2dv013.a2/protos"
	"google.golang.org/grpc"
)

// Represents the handler for logout route
type LogoutHandler struct {
	cookieService interfaces.ICookieService
	sessionLayer  pb.SessionLayerClient
}

// Constructor function
func NewLogoutHandler(cookieService interfaces.ICookieService, sessionLayer pb.SessionLayerClient) *LogoutHandler {
	return &LogoutHandler{
		cookieService: cookieService,
		sessionLayer:  sessionLayer,
	}
}

func (l *LogoutHandler) Handle(ctx *fiber.Ctx) error {
	log.Printf("Handling logout request")

	sessionToken := l.cookieService.GetSessionTokenFromCookie(ctx)
	if sessionToken == "" {
		log.Printf("No session token found in cookie")
		l.cookieService.SetDeletionCookie(ctx)
		return ctx.Status(fiber.StatusNoContent).SendString("Allready logged out")
	}

	var opts []grpc.CallOption
	removeSessionReq := pb.DeleteSessionRequest{SessionToken: sessionToken}
	removeSessionRes, err := l.sessionLayer.DeleteSession(ctx.Context(), &removeSessionReq, opts...)

	l.cookieService.SetDeletionCookie(ctx)

	if err != nil {
		log.Printf("Delete session gRPC call returned error: %s", err.Error())
		return ctx.Status(fiber.StatusNoContent).SendString("Session cleared")
	}

	switch removeSessionRes.Status {
	case pb.SessionStatus_Session_Success:
		return ctx.Status(fiber.StatusNoContent).SendString("Logged out successfully")
	case pb.SessionStatus_Session_No_Session_Found:
		return ctx.Status(fiber.StatusNoContent).SendString("Allready logged out")
	default:
		return ctx.Status(fiber.StatusNoContent).SendString("Session cleared")
	}
}

// Helper method for returning internalServerError
func (l *LogoutHandler) internalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong internally, try again later",
	})
}
