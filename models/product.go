package models

type Product struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Price       float32
	Category    Category
}
