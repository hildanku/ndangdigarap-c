package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hildanku/ndangdigarap/config"
	"github.com/hildanku/ndangdigarap/handlers"
	"github.com/hildanku/ndangdigarap/middlewares"
	"github.com/hildanku/ndangdigarap/utils"
)

func main() {

	// https://stackoverflow.com/questions/54456186/how-to-fix-environment-variables-not-working-while-running-from-system-d-service
	utils.LoadEnv()

	db := config.ConnectDatabase()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("init db failed", err)
	}
	defer sqlDB.Close()

	log.Println("Up!")

	app := fiber.New()
	app.Post("/register", handlers.RegisterUser(db))
	app.Post("/login", handlers.LoginUser(db, os.Getenv("JWT_SECRET")))
	app.Get("/healthcheck", handlers.Hello)

	api := app.Group("/api", middlewares.JWTMiddleware)
	api.Get("/protected", handlers.ProtectedEndpoint)

	log.Fatal(app.Listen(":3000"))

}
