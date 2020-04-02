package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment"`
	Name string `gorm:"not null"`
}

func (category *Category) FindAllCategories(db *gorm.DB) ([]*Category, error) {
	categories := make([]*Category, 0)

	if err := db.Table("product_category").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
