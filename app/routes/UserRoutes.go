package routes

import (
	"authMicroservice/app/http/controllers/userController"
	"authMicroservice/app/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	v1 := app.Group("/v1")
	v1.Post("/login", userController.Login)
	v1.Post("/registration", userController.CreateUser)
	v1.Get("/user", middlewares.AuthMiddleware, userController.GetUser)
	v1.Get("/logout", middlewares.AuthMiddleware, userController.Logout)
}
