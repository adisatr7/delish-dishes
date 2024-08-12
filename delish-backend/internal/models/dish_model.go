package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID        string 		`gorm:"primaryKey"`
	Name      string
	Desc      string
	DishOrder []DishOrder 	`gorm:"foreignKey:DishID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
