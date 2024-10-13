package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Status     string `gorm:"default:'active'"` // "active", "cancelled", "finished"
	UserID     uint   `gorm:"not null"`
	ShowtimeID uint   `gorm:"not null"`
	Seats      []Seat `gorm:"many2many:reservation_seats;"`
}
