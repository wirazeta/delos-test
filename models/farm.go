package models

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey" json:"ID"`
	Name      string         `gorm:"unique" json:"Name"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt"`
	Pond      []Pond
}

type CreateFarmInput struct {
	Name string `json:"Name" binding:"required"`
}

type UpdateFarmInput struct {
	Name string `json:"Name" binding:"required"`
}
