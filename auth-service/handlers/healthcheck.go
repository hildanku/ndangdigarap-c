package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hildanku/ndangdigarap/utils"
	"gorm.io/gorm"
)

func Hello(c *fiber.Ctx) error {
	null := 0
	return utils.AppResponse(c, fiber.StatusOK, "up!", null)
}

func HealthCheck(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			return utils.AppResponse(
				c,
				fiber.StatusInternalServerError,
				"Database is down",
				nil,
			)
		}
		return utils.AppResponse(
			c,
			fiber.StatusOK,
			"Database is up",
			nil,
		)
	}
}
