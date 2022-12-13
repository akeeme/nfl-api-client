package database

import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func getENV(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

func Connect() {
	db, err := gorm.Open(mysql.Open(getENV("dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

	DB = db

}