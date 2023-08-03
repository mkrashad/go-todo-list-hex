package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mkrashad/go-todo/user/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnvVariables() {
	err := godotenv.Load("/home/rashad/workspace/github.com/mkrashad/go-todo/user/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {
	var err error
	DbConfig := struct {
		Host     string
		User     string
		Password string
		DbName   string
		Port     string
	}{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbConfig.Host, DbConfig.User, DbConfig.Password, DbConfig.DbName, DbConfig.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&internal.User{})
	if err != nil {
		log.Fatal("Could not migrate:", err)
	}
	fmt.Println("Database migrated succesfully")
}
