package controllers

import (
	"catalog/api/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	databaseName := os.Getenv("databaseName")
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, databaseName)

	if s.DB, err = gorm.Open("mysql", url); err != nil {
		log.Printf("Error connecting to %s database", databaseName)
		log.Fatal(err)
	}

	s.DB.AutoMigrate(&models.Category{}, &models.Product{})
	s.DB.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)",
		"RESTRICT", "RESTRICT")

	s.Router = mux.NewRouter()
	s.initRoutes()
}

func (s *Server) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}

func (s *Server) CloseDB() {
	if err := s.DB.Close(); err != nil {
		log.Panic(err)
	}
}
