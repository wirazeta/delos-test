package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"ID"`
	Name      string `gorm:"unique" json:"Name"`
	Count     int    `json:"Count"`
	UserAgent int    `json:"UserAgent"`
}
