package main

import (
	"os"

	"github.com/vnkot/piklnk/internal/auth/repository"
	"github.com/vnkot/piklnk/internal/link"
	"github.com/vnkot/piklnk/internal/stat"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &repository.UserModel{}, &stat.Stat{})
}
