package models

import (
	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	FileId string `gorm:"size:100;unique;not null" json:"file_id"`
	UserId string `gorm:"size:50;not null" json:"user_id"`
	Path   string `gorm:"size:500;not null" json:"path"`
}

type UploadFile struct {
	UserId string `json:"userId" binding:"required"`
}

type DownloadFile struct {
	FileId string `json:"fileId" binding:"required"`
}

type RegFile struct {
	FileId string `json:"fileId"`
	UserId string `json:"userId"`
	Path   string `json:"path"`
}
