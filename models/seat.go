package models

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	Row    string `gorm:"not null"`
	Number int    `gorm:"not null"`
	HallID uint   `gorm:"not null"`
}
