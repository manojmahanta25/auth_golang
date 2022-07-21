package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	fmt.Print(err)
	code := fiber.StatusInternalServerError
	// Retrieve the custom status code if it's a fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if err != nil {
		return ctx.Status(code).JSON(fiber.Map{
			"ok":            false,
			"error":         err,
			"error_message": "Internal Server Error",
		})
	}
	return nil
}
