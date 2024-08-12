package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string 		`gorm:"primaryKey;type:uuid"`
	CustomerName string			`gorm:"type:varchar(256);index;not null"`
	TotalPrice   uint64			`gorm:"not null"`
	Note         *string
	DishOrders   []DishOrder 	`gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time		`gorm:"autoCreateTime"`
	UpdatedAt    time.Time		`gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt
}
