package api

import (
	"go-gin-file/app/models"
	"go-gin-file/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(id string) models.Users {
	var user models.Users
	if err := db.Where("user_id=?", id).First(&user); err != nil {
		return models.Users{}
	}
	return user
}

func getAllUser(c *gin.Context) {
	var users []models.Users
	if err := db.Find(&users); err != nil {
		common.Res(c, http.StatusInternalServerError, "Upload File Success")
		return
	}
	common.Res(c, http.StatusOK, users)
}

func registerUser(c *gin.Context) {
	p := new(models.RegUserParam)
	if err := c.Bind(&p); err != nil {
		log.Printf("Validation error: %s", err)
		common.Res(c, http.StatusBadRequest, err)
		return
	}
	users := models.Users{
		UserId: p.UserId,
		Name:   p.Name,
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Printf("Register users error : %s", result.Error)
		common.Res(c, http.StatusInternalServerError, result.Error)
		return
	}

	r := "Register User Success"

	common.Res(c, http.StatusOK, r)
}
