package models

import (
	"gorm.io/gorm"
)

type UserAgent struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"ID"`
	Name  string `json:"Name"`
	LogID uint   `json:"LogID"`
	Log   Log    `json:"Log"`
}
