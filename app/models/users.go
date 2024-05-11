package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserId string `gorm:"size:50;unique;not null" json:"user_id"`
	Name   string `gorm:"size:50;not null" json:"name"`
}

type RegUserParam struct {
	UserId string `json:"userId" binding:"required"`
	Name   string `json:"name" binding:"required"`
}
