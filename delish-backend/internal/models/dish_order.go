package models

import "gorm.io/gorm"

type DishOrder struct {
	gorm.Model
	ID      string `gorm:"primaryKey" json:"id"`
	DishID  string `gorm:"type:uuid;foreignKey" json:"dishId"`
	Dish    Dish   `gorm:"joinForeignKey:DishID" json:"dish"`
	OrderID string `gorm:"type:uuid;foreignKey" json:"orderId"`
	Order   Order  `gorm:"joinForeignKey:OrderID" json:"order"`
	Qty     uint64 `gorm:"not null" json:"qty"`
}
