package controllers

import (
	"catalog/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) InitDB() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	databaseName := os.Getenv("databaseName")
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, databaseName)

	if server.DB, err = gorm.Open("mysql", url); err != nil {
		log.Printf("Error connecting to %s database", databaseName)
		log.Fatal(err)
	}

	server.DB.AutoMigrate(&models.Category{}, &models.Product{})
	server.DB.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)",
		"RESTRICT", "RESTRICT")
}

func (server *Server) CloseDB() {
	if err := server.DB.Close(); err != nil {
		log.Panic(err)
	}
}
