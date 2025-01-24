package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hildanku/ndangdigarap/utils"
)

func ValidateToken(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Token is required", nil)
		}

		// Parse and validate token
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !parsedToken.Valid {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid token", nil)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}

		return utils.AppResponse(c, fiber.StatusOK, "Token is valid", fiber.Map{
			"user_id": claims["user_id"],
		})
	}
}
