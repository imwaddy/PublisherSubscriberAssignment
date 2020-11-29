package models

import "github.com/jinzhu/gorm"

type RatePlan struct {
	gorm.Model
	ID       string `json:"rate_plan_id";gorm:"primaryKey"`
	HotelID  string `json:"hotel_id"`
	Hotel    Hotel  `gorm:"foreignKey:ID;references:HotelID"`
	Name     string `json:"name"`
	MealPlan string `json:"meal_plan"`
	//CancellationPolicy CancellationPolicy `json:"meal_plan"`
	// OtherConditions    []string
}

type CancellationPolicy struct {
	Type              string
	ExpiresDaysBefore int
}
