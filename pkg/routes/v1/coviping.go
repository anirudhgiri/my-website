package routes

import "github.com/gofiber/fiber/v2"

func CoviPingRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/html/coviping/index.html")
	})
}
