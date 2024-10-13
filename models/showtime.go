package models

import (
	"time"

	"gorm.io/gorm"
)

type Showtime struct {
	gorm.Model
	StartAt      time.Time `gorm:"not null"`
	MovieID      uint      `gorm:"not null"`
	HallID       uint      `gorm:"not null"`
	Reservations []Reservation
}
