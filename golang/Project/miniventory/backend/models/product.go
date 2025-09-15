package models

import "gorm.io/gorm"

type Product struct {
	ID		uint `json:"id" gorm:"primaryKey"`
	Name	string `json:"name" binding:"required"`
	Price	int `json:"price" binding:"required,gte=0"`
	Stock	int `json:"stock" binding:"required,gte=0"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}