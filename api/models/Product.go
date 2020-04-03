package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint     `gorm:"primary_key;auto_increment" json:"id"`
	Name        string   `gorm:"not null" json:"name"`
	Description string   `gorm:"not null" json:"description"`
	Price       float32  `gorm:"not null" json:"price"`
	CategoryID  uint     `gorm:"foreignkey:category_id;not null" json:"category_id"`
	Category    Category `json:"category"`
}

func (p *Product) Save(db *gorm.DB) (*Product, error) {
	if err := db.Save(&p).Error; err != nil {
		return nil, err
	}

	if err := db.Take(&p.Category, "id = ?", p.CategoryID).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Update(db *gorm.DB, id uint) (*Product, error) {
	if err := db.Model(&p).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"price":       p.Price,
		"category_id": p.CategoryID,
	}).Error; err != nil {
		return nil, err
	}

	if err := db.Take(&p.Category, "id = ?", p.CategoryID).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Delete(db *gorm.DB, id uint) error {
	if err := db.Delete(&Product{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) FindAll(db *gorm.DB) ([]*Product, error) {
	ps := make([]*Product, 0)
	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}

	if len(ps) > 0 {
		for _, p := range ps {
			if err := db.Take(&p.Category, "id = ?", p.CategoryID).Error; err != nil {
				return nil, err
			}
		}
	}

	return ps, nil
}

func (p *Product) FindAllByCategoryId(db *gorm.DB, id uint) ([]*Product, error) {
	c, err := (&Category{}).FindById(db, id)
	if err != nil {
		return nil, err
	}

	ps := make([]*Product, 0)
	if err := db.Model(&c).Related(&ps).Error; err != nil {
		return nil, err
	}

	if len(ps) > 0 {
		for _, p := range ps {
			if err := db.Take(&p.Category, "id = ?", p.CategoryID).Error; err != nil {
				return nil, err
			}
		}
	}

	return ps, nil
}
