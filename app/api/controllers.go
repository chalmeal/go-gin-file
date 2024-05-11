package api

import (
	"go-gin-file/app/models"
	"go-gin-file/common"

	"github.com/gin-gonic/gin"
)

var (
	db = common.DbConnect()
)

func setUp() {
	db.AutoMigrate(
		&models.Files{},
		&models.Users{},
	)
}

func RegisterRouter(r *gin.RouterGroup) {
	// setup db auto migrate
	{
		setUp()
	}

	{
		// User
		r.GET("/user", getAllUser)
		r.POST("/user", registerUser)

		// File
		r.PUT("/file/upload", fileUpload)
		r.POST("/file/download", getFilePath)
	}
}
