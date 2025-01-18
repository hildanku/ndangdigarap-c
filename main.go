package main

import (
	"log"

	"github.com/hildanku/ndangdigarap/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
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

}
