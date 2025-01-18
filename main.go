package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hildanku/ndangdigarap/config"
	"github.com/hildanku/ndangdigarap/handlers"
	"github.com/hildanku/ndangdigarap/middlewares"
	"github.com/joho/godotenv"
)

func main() {

	// https://stackoverflow.com/questions/54456186/how-to-fix-environment-variables-not-working-while-running-from-system-d-service
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println("pwd", pwd)

	//use ../.env because main.go inside /cmd
	err = godotenv.Load(filepath.Join(pwd, "./.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("JWT Secret:", os.Getenv("JWT_ACCESS_SECRET"))

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// jwtSecret := os.Getenv("JWT_ACCESS_SECRET")
	// if jwtSecret == "" {
	// 	log.Fatal("JWT_SSECRET is not set in .env")
	// }
	// log.Println("JWT Secret:", jwtSecret)

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
