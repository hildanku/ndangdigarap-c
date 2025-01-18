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

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString([]byte(jwtSecret))
		log.Println("access", jwtSecret)
		if err != nil {
			return utils.AppResponse(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
		}

		return utils.AppResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
			"token": tokenString,
		})
	}
}

func ProtectedEndpoint(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return c.JSON(fiber.Map{
		"message":  "This is a protected endpoint",
		"username": username,
	})
}
