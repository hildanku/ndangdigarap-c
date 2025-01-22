package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hildanku/ndangdigarap/utils"
)

func Hello(c *fiber.Ctx) error {
	null := 0
	return utils.AppResponse(c, fiber.StatusOK, "up!", null)
}
