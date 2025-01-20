package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Represents a handler
type Handler interface {
	Handle(ctx *fiber.Ctx) error
}

// Represents the auth middleware
type AuthMiddleware interface {
	Authenticate(ctx *fiber.Ctx) error
}

// Enum type
type HandlerType int

// Handler enums
const (
	CreateUserHandlerInstance HandlerType = iota
	LoginHandlerInstance
	LogoutHandlerInstance
	ValidateSessionHandlerInstance
	RefreshSessionHandlerInstance
	CreateTaskHandlerInstance
	DeleteTaskHandlerInstance
	UpdateTaskHandlerInstance
	GetTaskHandlerInstance
	GetMultipleTasksHandlerInstance
)

// Represents a factory for creating handlers and auth middleware
type HandlerFactory interface {
	CreateHandler(handler HandlerType) Handler
	CreateAuthMiddleware() AuthMiddleware
}

// Represents the router
type Router struct {
	port        string
	hostBaseUrl string
	factory     HandlerFactory
}

// Constructor function
func NewRouter(port string, hostBaseUrl string, factory HandlerFactory) *Router {
	return &Router{
		port:        port,
		hostBaseUrl: hostBaseUrl,
		factory:     factory,
	}
}

// Function for starting the router
func (r *Router) Start() {
	app := fiber.New()

	// Add cors settings
	corsConfig := cors.Config{
		AllowOrigins:     r.hostBaseUrl,
		AllowCredentials: true,
		MaxAge:           3600,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}
	app.Use(cors.New(corsConfig))

	authMiddleware := r.factory.CreateAuthMiddleware()

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")

	createUserHandler := r.factory.CreateHandler(CreateUserHandlerInstance)
	users.Post("/", createUserHandler.Handle)
	loginHandler := r.factory.CreateHandler(LoginHandlerInstance)
	users.Post("/login", loginHandler.Handle)

	usersLogout := users.Group("/logout")
	logoutHandler := r.factory.CreateHandler(LogoutHandlerInstance)
	usersLogout.Delete("/", logoutHandler.Handle)

	auth := v1.Group("/auth")
	validateSessionHandler := r.factory.CreateHandler(ValidateSessionHandlerInstance)
	auth.Get("/session", validateSessionHandler.Handle)
	refreshSessionHandler := r.factory.CreateHandler(RefreshSessionHandlerInstance)
	auth.Post("/session/refresh", refreshSessionHandler.Handle)

	userTasks := users.Group("/me/tasks")
	userTasks.Use(authMiddleware.Authenticate)
	getTaskHandler := r.factory.CreateHandler(GetTaskHandlerInstance)
	userTasks.Get("/:taskId", getTaskHandler.Handle)
	getMultipleTasksHandler := r.factory.CreateHandler(GetMultipleTasksHandlerInstance)
	userTasks.Get("/", getMultipleTasksHandler.Handle)
	createTaskHandler := r.factory.CreateHandler(CreateTaskHandlerInstance)
	userTasks.Post("/", createTaskHandler.Handle)
	updateTaskHandler := r.factory.CreateHandler(UpdateTaskHandlerInstance)
	userTasks.Put("/:taskId", updateTaskHandler.Handle)
	deleteTaskHandler := r.factory.CreateHandler(DeleteTaskHandlerInstance)
	userTasks.Delete("/:taskId", deleteTaskHandler.Handle)

	log.Fatal(app.Listen(r.port))
}
