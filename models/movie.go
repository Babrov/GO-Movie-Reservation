package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	PosterImage string
	Genre       string
	Duration    *time.Duration
	Showtimes   []Showtime
}
