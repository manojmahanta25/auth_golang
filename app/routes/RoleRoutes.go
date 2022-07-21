package routes

import (
	"authMicroservice/app/http/controllers/roleController"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app fiber.Router) {
	v1 := app.Group("/v1")
	v1.Post("/roles", roleController.CreateRole)
}
