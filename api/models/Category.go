package models

import (
	"catalog/api/utils"
	"github.com/jinzhu/gorm"
	"strings"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (c *Category) validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return utils.ErrInvalidCategoryName
	}

	return nil
}

func (c *Category) Save(db *gorm.DB) (*Category, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) Update(db *gorm.DB, id uint) (*Category, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	if err := db.Model(&c).Where("id = ?", id).Updates(map[string]interface{}{
		"name": c.Name,
	}).Error; err != nil {
		return nil, err
	}

	if err := db.Take(&c, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, utils.ErrCategoryNotFound
		}
		return nil, err
	}

	return c, nil
}

func (c *Category) Delete(db *gorm.DB, id uint) error {
	if d := db.Delete(&Category{}, "id = ?", id); d.RowsAffected == 0 {
		return utils.ErrCategoryDeletionFailed
	}

	return nil
}

func (c *Category) FindById(db *gorm.DB, id uint) (*Category, error) {
	if err := db.Take(&c, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, utils.ErrCategoryNotFound
		}
		return nil, err
	}

	return c, nil
}

func (c *Category) FindAll(db *gorm.DB) ([]*Category, error) {
	cs := make([]*Category, 0)
	if err := db.Find(&cs).Error; err != nil {
		return nil, err
	}

	return cs, nil
}
