package modeltests

import (
	"catalog/api/controllers"
	"catalog/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var (
	server   = controllers.Server{}
	category = models.Category{}
	product  = models.Product{}
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(os.ExpandEnv("../../.env")); err != nil {
		log.Fatalf("fail getting env %v\n", err)
	}
	SetupDatabase()
	os.Exit(m.Run())
}

func SetupDatabase() {
	testUsername := os.Getenv("test_username")
	testPassword := os.Getenv("test_password")
	testDatabaseName := os.Getenv("test_databaseName")
	testUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", testUsername, testPassword, testDatabaseName)

	var err error
	server.DB, err = gorm.Open("mysql", testUrl)
	if err != nil {
		log.Printf("Error connecting to %s database", testDatabaseName)
		log.Fatal(err)
	}
}

func refreshCategoryTable() error {
	if err := server.DB.DropTableIfExists(&models.Category{}).Error; err != nil {
		return err
	}

	if err := server.DB.AutoMigrate(&models.Category{}).Error; err != nil {
		return err
	}

	return nil
}

func refreshProductAndCategoryTables() error {
	if err := server.DB.DropTableIfExists(&models.Product{}).Error; err != nil {
		return err
	}

	if err := server.DB.DropTableIfExists(&models.Category{}).Error; err != nil {
		return err
	}

	if err := server.DB.AutoMigrate(&models.Category{}, &models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func seedCategories() error {
	cs := []models.Category{
		{Name: "Laptops"},
		{Name: "Phones"},
	}

	for _, c := range cs {
		if err := server.DB.Model(&models.Category{}).Create(&c).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedCategoryAndProducts() error {
	cs := []models.Category{
		{Name: "Electronics"},
		{Name: "Food"},
	}

	ps := []models.Product{
		{Name: "Mouse", Description: "Logitech", Price: 50.99, CategoryID: 1},
		{Name: "Keyboard", Description: "Keychrone", Price: 75, CategoryID: 1},
		{Name: "Water", Description: "Coconut water", Price: 3.75, CategoryID: 2},
	}

	for _, c := range cs {
		if err := server.DB.Model(&models.Category{}).Create(&c).Error; err != nil {
			return err
		}
	}

	for _, p := range ps {
		if err := server.DB.Model(&models.Category{}).Create(&p).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedSingleCategory() (*models.Category, error) {
	c := &models.Category{Name: "Electronics"}
	if err := server.DB.Model(&models.Category{}).Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
