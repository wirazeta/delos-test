package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey" json:"ID"`
	Name      string         `gorm:"unique" json:"Name"`
	Count     int            `json:"Count"`
	UserAgent string         `json:"UserAgent"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt"`
}
