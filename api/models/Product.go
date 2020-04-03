package models

import "github.com/jinzhu/gorm"

type Product struct {
	ID          uint    `gorm:"primary_key;auto_increment"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float32 `gorm:"not null"`
	CategoryID  uint    `gorm:"foreignkey:category_id"`
	Category    Category
}

func (p *Product) Save(db *gorm.DB) (*Product, error) {
	if err := db.Save(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Update(db *gorm.DB, id uint) (*Product, error) {
	if err := db.Model(&p).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"price":       p.Price,
		"category_id": p.Category.ID,
	}).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).Take(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) FindAll(db *gorm.DB) ([]*Product, error) {
	ps := make([]*Product, 0)
	if err := db.Find(&ps).Error; err != nil {
		return nil, err
	}

	return ps, nil
}

func (p *Product) FindAllByCategory(db *gorm.DB, category Category) ([]*Product, error) {
	ps := make([]*Product, 0)
	if err := db.Model(&category).Related(&ps).Error; err != nil {
		return nil, err
	}

	return ps, nil
}
