package interfaces

import "github.com/gofiber/fiber/v2"

// Interface for cookie service
type ICookieService interface {
	GetUserIdFromCookie(ctx *fiber.Ctx) string
	GetSessionTokenFromCookie(ctx *fiber.Ctx) string
	SetDeletionCookie(ctx *fiber.Ctx)
	SetSessionCookie(ctx *fiber.Ctx, token string, userId string)
}
