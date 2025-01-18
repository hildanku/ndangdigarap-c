package config

import (
	"os"
	"testing"
	// "gorm.io/gorm"
)

func TestConnectDatabaseSuccess(t *testing.T) {

	os.Setenv("MYSQL_DATABASE", "xx")
	os.Setenv("MYSQL_USER", "x")
	os.Setenv("MYSQL_PASSWORD", "x")
	os.Setenv("MYSQL_HOST", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")

	db := ConnectDatabase()
	if db == nil {
		t.Fatal("Expected a valid database connection, got nil")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to retrieve SQL DB: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Database ping failed: %v", err)
	}
}
