package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Represents a cookie service
type CookieService struct {
	sessionCookieName string
	ttl               time.Duration
}

// Constructor function
func NewCookieService(sessionCookieName string, ttl time.Duration) *CookieService {
	return &CookieService{
		sessionCookieName: sessionCookieName,
		ttl:               ttl,
	}
}

// Function for extracting the user id from a cookie
func (c *CookieService) GetUserIdFromCookie(ctx *fiber.Ctx) string {
	cookieVal := ctx.Cookies(c.sessionCookieName)
	if cookieVal == "" {
		return ""
	}

	splitted := strings.Split(cookieVal, ":")
	return splitted[0]
}

// Function for extracting the session token from a cookie
func (c *CookieService) GetSessionTokenFromCookie(ctx *fiber.Ctx) string {
	cookieVal := ctx.Cookies(c.sessionCookieName)
	if cookieVal == "" {
		return ""
	}

	splitted := strings.Split(cookieVal, ":")
	return splitted[1]
}

// Function for setting a deletion token
func (c *CookieService) SetDeletionCookie(ctx *fiber.Ctx) {
	ctx.Cookie(&fiber.Cookie{
		Name:     c.sessionCookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   true,
	})
}

// Function for setting a session token
func (c *CookieService) SetSessionCookie(ctx *fiber.Ctx, token string, userId string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     c.sessionCookieName,
		Value:    fmt.Sprintf("%s:%s", userId, token),
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
}
