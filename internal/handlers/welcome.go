package handlers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	return c.Render("welcome", nil, "layouts/main") // redirects to main page in html when welcome function is called
}
