package routes

import "github.com/gofiber/fiber/v2"

func RootRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/html/homepage/index.html")
	})

	router.Route("/coviping", CoviPingRouter)
}
