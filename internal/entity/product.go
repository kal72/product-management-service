package entity

import "time"

type Product struct {
	ID         int     `gorm:"primaryKey"`
	Name       string  `gorm:"size:255;not null"`
	Price      float64 `gorm:"type:decimal(12,2);not null"`
	Stock      int     `gorm:"default:0"`
	CategoryID int     `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type ProductDetail struct {
	ID           int
	Name         string
	Price        float64
	Stock        int
	CategoryID   int
	CategoryName string
	CreatedAt    time.Time
}
