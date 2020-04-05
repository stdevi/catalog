package models

import (
	"catalog/api/utils"
	"github.com/jinzhu/gorm"
	"strings"
)

type Product struct {
	ID          uint     `gorm:"primary_key;auto_increment" json:"id"`
	Name        string   `gorm:"not null" json:"name"`
	Description string   `json:"description"`
	Price       float32  `gorm:"not null" json:"price"`
	CategoryID  uint     `gorm:"foreignkey:category_id;not null" json:"category_id"`
	Category    Category `json:"category"`
}

func (p *Product) validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return utils.ErrInvalidProductName
	}
	if p.Price < 0 {
		return utils.ErrInvalidProductPrice
	}

	return nil
}

func (p *Product) Save(db *gorm.DB) (*Product, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}

	if err := db.Save(&p).Error; err != nil {
		return nil, err
	}

	if err := db.Take(&p.Category, "id = ?", p.CategoryID).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Update(db *gorm.DB, id uint) (*Product, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}

	c, err := (&Category{}).FindById(db, p.CategoryID)
	if err != nil {
		return nil, err
	}
	p.Category = *c

	if err := db.Model(&p).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"price":       p.Price,
		"category_id": p.CategoryID,
	}).Error; err != nil {
		return nil, err
	}

	if err := db.Take(&p, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, utils.ErrProductNotFound
		}
		return nil, err
	}

	return p, nil
}

func (p *Product) Delete(db *gorm.DB, id uint) error {
	if d := db.Delete(&Product{}, "id = ?", id); d.RowsAffected == 0 {
		return utils.ErrProductDeletionFailed
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
