package routes

import (
	"my-website/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func BlogRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		blogs, _ := db.GetAllBlogs()
		return c.Render("blog_list", fiber.Map{"Title": "Anirudh Giri - Blog", "Viewname": "blog_list", "Blog": true, "Blogs": blogs})
	})

	router.Get("/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		blog := db.GetBlogBySlug(slug)
		mdParser := parser.NewWithExtensions(parser.HardLineBreak | parser.FencedCode | parser.Strikethrough | parser.Attributes)
		contentHTML := markdown.ToHTML([]byte(blog.Content), mdParser, nil)
		return c.Render("blog_content", fiber.Map{"blogContent": string(contentHTML), "Blog": true, "Viewname": "blog_content"}, "layouts/blog_content")
	})
}
