package entity

import "time"

type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
