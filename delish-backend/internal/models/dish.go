package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	ID        string         `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string         `gorm:"type:varchar(256);index;not null" json:"name"`
	Desc      *string        `gorm:"type:varchar(256)" json:"desc"`
	Price     uint64         `gorm:"default:0;not null" json:"price"`
	DishOrder []DishOrder    `gorm:"foreignKey:DishID" json:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
