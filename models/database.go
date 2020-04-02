package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	db = openConn()
}

func openConn() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	databaseName := os.Getenv("databaseName")
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, databaseName)

	db, err := gorm.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
