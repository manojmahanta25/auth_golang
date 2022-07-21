package app

import (
	"authMicroservice/app/model"
	"authMicroservice/app/routes"
	"authMicroservice/app/utils/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋! Welcome to Auth Server. We are here for your authentication")
}

func headers(c *fiber.Ctx) error {
	// for cors
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Set("Access-Control-Allow-Methods", "PUT, PATCH, POST, GET, DELETE, OPTIONS")
	r := string(c.Request().Header.Method())
	if r == "OPTIONS" {
		c.Set("Access-Control-Allow-Methods", "PUT, PATCH, POST, GET, DELETE, OPTIONS")
		return c.Status(200).JSON(fiber.Map{})

	}
	return c.Next()
}

func App(appName string) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		ServerHeader:  "Authentication",
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       appName,
		ErrorHandler:  handlers.DefaultErrorHandler,
	})
	model.Migrate()
	app.Use(recover.New())
	app.Use(headers)
	routes.Routes(app)
	app.Get("/", helloWorld)
	return app
}
