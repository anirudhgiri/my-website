package main

import (
	"log"
	"my-website/pkg/constants"
	"my-website/pkg/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(logger.New())
	app.Static("/", "./static")
	app.Route("/", routes.RootRouter)
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Redirect("/")
	})
	log.Fatalln(app.Listen(constants.PORT))
}
