package models

import (
	"time"

	"gorm.io/gorm"
)

type Pond struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey" json:"ID"`
	Name      string         `gorm:"unique" json:"Name"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt"`
	FarmID    int
	Farm      Farm
}

type CreatePond struct {
	FarmID int    `json:"FarmID" binding:"required"`
	Name   string `json:"Name" binding:"required"`
}

type UpdatePond struct {
	FarmID int    `json:"FarmID"`
	Name   string `json:"Name" binding:"required"`
}
