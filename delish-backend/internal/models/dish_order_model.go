package models

import "gorm.io/gorm"

type DishOrder struct {
    gorm.Model
    ID      string `gorm:"primaryKey"`
    DishID  string `gorm:"not null"`
    OrderID string `gorm:"not null"`
    Dish    Dish   `gorm:"foreignKey:DishID;references:ID"`
    Order   Order  `gorm:"foreignKey:OrderID;references:ID"`
}