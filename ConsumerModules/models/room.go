package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	ID          string `json:"room_id", gorm:"primary_key"`
	HotelID     string `json:"hotel_id"`
	Hotel       Hotel  `gorm:"foreignKey:ID;references:HotelID"`
	Description string `json:"description"`
	Name        string `json:"name"`
	//	Capacity    capacity
}

type capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}
