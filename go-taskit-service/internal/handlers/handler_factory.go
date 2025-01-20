package handlers

import (
	"log"

	"gn222gq.2dv013.a2/internal"
	"gn222gq.2dv013.a2/internal/interfaces"
	"gn222gq.2dv013.a2/internal/middleware"
	pb "gn222gq.2dv013.a2/protos"
)

// Represents a handler factory for creating handler instances
type HandlerFactory struct {
	dataLayer     pb.DataLayerClient
	sessionLayer  pb.SessionLayerClient
	cookieService interfaces.ICookieService
	authService   interfaces.IAuthService
	messageBroker interfaces.IMessageBroker
}

// Constructor function
func NewHandlerFactory(dataLayer pb.DataLayerClient,
	sessionLayer pb.SessionLayerClient,
	cookieService interfaces.ICookieService,
	authService interfaces.IAuthService,
	messageBroker interfaces.IMessageBroker) *HandlerFactory {
	return &HandlerFactory{
		dataLayer:     dataLayer,
		sessionLayer:  sessionLayer,
		cookieService: cookieService,
		authService:   authService,
		messageBroker: messageBroker,
	}
}

// Function for creating a handler
func (h *HandlerFactory) CreateHandler(handler internal.HandlerType) internal.Handler {
	switch handler {
	case internal.CreateUserHandlerInstance:
		return NewCreateUserHandler(h.authService, h.dataLayer)
	case internal.LoginHandlerInstance:
		return NewLoginHandler(h.dataLayer, h.sessionLayer, h.authService, h.cookieService)
	case internal.LogoutHandlerInstance:
		return NewLogoutHandler(h.cookieService, h.sessionLayer)
	case internal.ValidateSessionHandlerInstance:
		return NewValidateSessionHandler(h.sessionLayer, h.cookieService)
	case internal.RefreshSessionHandlerInstance:
		return NewRefreshSessionHandler(h.cookieService, h.sessionLayer)
	case internal.CreateTaskHandlerInstance:
		return NewCreateTaskHandler(h.dataLayer, h.messageBroker)
	case internal.DeleteTaskHandlerInstance:
		return NewDeleteTaskHandler(h.dataLayer)
	case internal.UpdateTaskHandlerInstance:
		return NewUpdateTaskHandler(h.dataLayer, h.messageBroker)
	case internal.GetTaskHandlerInstance:
		return NewGetTaskHandler(h.dataLayer)
	case internal.GetMultipleTasksHandlerInstance:
		return NewGetMultipleTasksHandler(h.dataLayer)
	default:
		log.Fatalf("Handler factory create handler recieved unexpected handler type integer")
		return nil
	}
}

// Function for creating auth middleware
func (h *HandlerFactory) CreateAuthMiddleware() internal.AuthMiddleware {
	return middleware.NewAuthMiddleware(h.dataLayer, h.sessionLayer, h.authService, h.cookieService)
}
