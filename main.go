package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hildanku/ndangdigarap/config"
	"github.com/hildanku/ndangdigarap/handlers"
	"github.com/hildanku/ndangdigarap/middlewares"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Err load env")
	}

	db := config.ConnectDatabase()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("init db failed", err)
	}
	defer sqlDB.Close()

	log.Println("Up!")

	app := fiber.New()
	app.Post("/register", handlers.RegisterUser(db))
	app.Post("/login", handlers.LoginUser(db))
	app.Get("/healthcheck", handlers.Hello)

	api := app.Group("/api", middlewares.JWTMiddleware)
	api.Get("/protected", handlers.ProtectedEndpoint)

	log.Fatal(app.Listen(":3000"))

}
