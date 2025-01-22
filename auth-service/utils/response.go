package utils

import "github.com/gofiber/fiber/v2"

func AppResponse(c *fiber.Ctx, status int, message string, results interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"results": results,
	})
}
