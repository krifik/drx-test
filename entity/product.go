package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex;"`
	Name        string `gorm:"size:256"`
	Quantity    int    `gorm:"size:256"`
	Price       int    `gorm:"size:256"`
	Description string `gorm:"size:256"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
