package middlewares

import (
	"authMicroservice/app/http/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func parseCookie(ctx *fiber.Ctx) string {
	return ctx.Cookies("Authorization")
}

func parseHeader(ctx *fiber.Ctx) string {
	return ctx.Get("Authorization")
}

func parseQuery(ctx *fiber.Ctx) string {
	return ctx.Query("authorization")
}

func grabToken(ctx *fiber.Ctx) string {
	token := parseCookie(ctx)
	if token == "" {
		token = parseHeader(ctx)
		if token == "" {
			token = parseQuery(ctx)
		}
	}
	return token
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := grabToken(ctx)
	if token == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"ok":            false,
			"error":         "Auth Fail",
			"error_message": "Invalid Auth Format/empty",
		})
	}
	sub, err := services.JwtVerify(token)
	if err != nil {
		fmt.Println("Error", err.Error())
		return ctx.Status(401).JSON(fiber.Map{
			"ok":            false,
			"error":         "Auth Fail",
			"error_message": "Invalid Auth token",
		})
	}
	ctx.Locals("userInfo", sub)
	return ctx.Next()
}
