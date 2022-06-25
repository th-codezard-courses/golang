package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU        string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Name       string `gorm:"not null"`
	Desc       string
	Price      float64
	Status     uint   `gorm:"not null"`
	Image      string `gorm:"not null"`
	CategoryID uint
	Category   Category `gorm:"not null"`
}
