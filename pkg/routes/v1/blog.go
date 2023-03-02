package routes

import (
	"math"
	"my-website/pkg/db"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func BlogRouter(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		blogs, _ := db.GetAllBlogs()

		return c.Render("blog_list", fiber.Map{
			"Title":    "Anirudh Giri - Blog",
			"Viewname": "blog_list",
			"Blog":     true,
			"Blogs":    blogs,
		})
	})

	router.Get("/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		blog := db.GetBlogBySlug(slug)

		mdParser := parser.NewWithExtensions(parser.HardLineBreak | parser.FencedCode | parser.Strikethrough | parser.Attributes)
		contentHTML := markdown.ToHTML([]byte(blog.Content), mdParser, nil)
		readTime := math.Max(1, math.Round(float64(len(strings.Split(string(contentHTML), "")))/250))

		return c.Render("blog_content", fiber.Map{
			"blogTitle":     blog.Title,
			"blogSubTitle":  blog.SubTitle,
			"blogCreatedAt": blog.CreatedAt.Format("Jan 02, 2006"),
			"blogContent":   string(contentHTML),
			"Blog":          true,
			"Viewname":      "blog_content",
			"blogReadTime":  readTime,
		},
			"layouts/blog_content")
	})
}
