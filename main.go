package main

import (
	"my-website/pkg/constants"
	"my-website/pkg/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Static("/", "./static")
	app.Route("/", routes.RootRouter)
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Error 404: Page Not Found")
	})
	app.Listen(constants.PORT)
}
