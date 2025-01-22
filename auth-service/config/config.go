package config

import (
	"log"
	"os"

	"github.com/hildanku/ndangdigarap/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("err load env")
	// }

	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := "3306"

	log.Println(mysqlPort)

	dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("fail connect db", err)
	}
	log.Println("success")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	return db
}
