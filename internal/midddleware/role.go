package middleware

import "github.com/gofiber/fiber/v2"

func AdminRoleMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied. Admin role required.",
		})
	}

	return c.Next()
}
