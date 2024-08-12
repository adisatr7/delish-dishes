package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	ID        string      `gorm:"primaryKey;type:uuid"`
	Name      string      `gorm:"type:varchar(256);index;not null"`
	Desc      *string     `gorm:"type:varchar(256)"`
	DishOrder []DishOrder `gorm:"foreignKey:DishID"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
