package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	apiRoutes := app.Group("/api")
	UserRoutes(apiRoutes)
	RoleRoutes(apiRoutes)
}
