package models

import "gorm.io/gorm"

type DishOrder struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	DishID  string `gorm:"type:uuid;foreignKey"`
    Dish    Dish   `gorm:"joinForeignKey:DishID"`
	OrderID string `gorm:"type:uuid;foreignKey"`
    Order   Order  `gorm:"joinForeignKey:OrderID"`
}
