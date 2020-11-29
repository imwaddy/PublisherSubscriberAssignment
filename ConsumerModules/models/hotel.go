package models

import "github.com/jinzhu/gorm"

type Hotel struct {
	gorm.Model
	ID          string  `json:"hotel_id";gorm:"type:longtext"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Telephone   string  `json:"telephone"`
	Description string  `json:"description"`
	RoomCount   int     `json:"room_count"`
	Currency    string  `json:"currency"`
	// Amenities   []string `json:"amenities"`
}
