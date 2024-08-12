package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string 		`gorm:"primaryKey"`
	CustomerName string
	TotalPrice   string
	Note         *string
	DishOrders   []DishOrder 	`gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
