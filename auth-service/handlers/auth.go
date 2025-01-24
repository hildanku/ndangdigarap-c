package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hildanku/ndangdigarap/models"
	"github.com/hildanku/ndangdigarap/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func RegisterUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)

		if err := c.BodyParser(user); err != nil {
			return utils.AppResponse(c, fiber.StatusBadRequest, "Invalid request payload", nil)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return utils.AppResponse(c, fiber.StatusInternalServerError, "Failed to hash password", nil)
		}

		user.PasswordHash = string(hashedPassword)
		if err := db.Create(user).Error; err != nil {
			return utils.AppResponse(c, fiber.StatusInternalServerError, "Failed to register user", nil)
		}

		return utils.AppResponse(c, fiber.StatusCreated, "User registered successfully", nil)
	}
}

func LoginUser(db *gorm.DB, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginRequest := new(LoginRequest)

		if err := c.BodyParser(loginRequest); err != nil {
			return utils.AppResponse(c, fiber.StatusBadRequest, "Invalid request payload", nil)
		}

		var user models.User
		if err := db.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid username or password", nil)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password)); err != nil {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid username or password", nil)
		}

		accessToken, err := utils.GenerateToken(user.ID)
		log.Println("LOG FROM AUTH", accessToken)
		if err != nil {
			return utils.AppResponse(c, fiber.StatusInternalServerError, "Failed to generate access token", nil)
		}

		refreshToken := utils.GenerateRefreshToken()
		expiresAt := time.Now().Add(7 * 24 * time.Hour)
		db.Create(&models.Token{
			Token:     refreshToken,
			UserID:    user.ID,
			ExpiresAt: expiresAt,
		})

		return utils.AppResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func ProtectedEndpoint(c *fiber.Ctx) error {
	userToken, ok := c.Locals("user").(*jwt.Token)
	if !ok || userToken == nil {
		return utils.AppResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	username, ok := claims["username"].(string)
	if !ok {
		return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid Token", nil)
	}

	return utils.AppResponse(c, fiber.StatusOK, "This is a protected endpoint", fiber.Map{
		"username": username,
	})
}

func RefreshToken(db *gorm.DB, jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.FormValue("refresh_token")
		if refreshToken == "" {
			return utils.AppResponse(c, fiber.StatusBadRequest, "Refresh token is required", nil)
		}

		var token models.Token
		if err := db.Where("token = ?", refreshToken).First(&token).Error; err != nil {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Invalid refresh token", nil)
		}

		if time.Now().After(token.ExpiresAt) {
			return utils.AppResponse(c, fiber.StatusUnauthorized, "Refresh token expired", nil)
		}

		accessToken, err := utils.GenerateToken(token.UserID)
		if err != nil {
			return utils.AppResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
		}

		return utils.AppResponse(c, fiber.StatusOK, "Token refreshed", fiber.Map{
			"access_token": accessToken,
		})
	}
}
