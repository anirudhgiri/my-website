package routes

import "github.com/gofiber/fiber/v2"

func RootRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("homepage", fiber.Map{"Title": "Anirudh Giri", "Viewname": "homepage", "Homepage": true})
	})

	router.Route("/coviping", CoviPingRouter)

	router.Route("/blog", BlogRouter)
}
