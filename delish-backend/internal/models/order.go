package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string         `gorm:"primaryKey;type:uuid" json:"id"`
	CustomerName string         `gorm:"type:varchar(256);index;not null" json:"customerName"`
	TotalPrice   uint64         `gorm:"not null" json:"totalPrice"`
	Note         *string        `gorm:"type:varchar(256)" json:"note"`
	DishOrders   []DishOrder    `gorm:"foreignKey:OrderID" json:"-"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt,omitempty"`
}
