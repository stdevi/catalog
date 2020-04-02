package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment"`
	Name string `gorm:"not null;unique"`
}

var tableName = "product_category"

func (category *Category) Save(db *gorm.DB) (*Category, error) {
	if err := db.Table(tableName).Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (category *Category) Update(db *gorm.DB, id uint) (*Category, error) {
	if err := db.Table(tableName).Where("id = ?", id).Updates(map[string]interface{}{
		"name": category.Name,
	}).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (category *Category) FindAll(db *gorm.DB) ([]*Category, error) {
	categories := make([]*Category, 0)

	if err := db.Table(tableName).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
