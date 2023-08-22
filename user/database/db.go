package database

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mkrashad/go-todo/user/ctxutils"
)

var DB *gorm.DB

func ConnectToDB(ctx context.Context) {
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
		ctxutils.GetRequestLogger(ctx).Fatal("Failed to connect to database!")
	}
}
