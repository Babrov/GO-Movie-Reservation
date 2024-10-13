package models

import (
	"gorm.io/gorm"
)

type Hall struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Capacity  int    `gorm:"not null"`
	Seats     []Seat
	Showtimes []Showtime
}
