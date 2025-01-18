package handlers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	null := 0
	return c.JSON(fiber.Map{
		"message": "API is UP!",
		"data":    null,
	})
}
