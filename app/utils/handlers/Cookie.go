package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetCookie(name string, value string, duration time.Duration, secure bool, httpOnly bool) *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(duration)
	cookie.Secure = secure
	cookie.HTTPOnly = httpOnly
	return cookie
}
