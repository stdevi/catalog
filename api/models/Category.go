package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment"`
	Name string `gorm:"not null;unique"`
}

func (c *Category) Save(db *gorm.DB) (*Category, error) {
	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) Update(db *gorm.DB, id uint) (*Category, error) {
	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"name": c.Name,
	}).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) Delete(db *gorm.DB, id uint) error {
	if err := db.Delete(&Category{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (c *Category) FindAll(db *gorm.DB) ([]*Category, error) {
	cs := make([]*Category, 0)

	if err := db.Find(&cs).Error; err != nil {
		return nil, err
	}

	return cs, nil
}
