package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorJsonOutput(c *fiber.Ctx, err error, status int, message ...string) error {
	if e, ok := err.(*fiber.Error); ok {
		status = e.Code
	}
	return c.Status(status).JSON(fiber.Map{
		"ok":            false,
		"error":         err.Error(),
		"error_message": message,
	})
}
func SuccessJsonOutputData(c *fiber.Ctx, data interface{}, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"ok":   true,
		"data": data,
	})

}

func JsonOutputOrError(c *fiber.Ctx, result interface{}, err error, successStatus int, errStatus int) error {
	if err != nil {
		fmt.Println("Error", err)
		return ErrorJsonOutput(c, err, errStatus)
	}
	return SuccessJsonOutputData(c, result, successStatus)
}
