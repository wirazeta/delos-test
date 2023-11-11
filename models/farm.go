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
	DeletedAt gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}
