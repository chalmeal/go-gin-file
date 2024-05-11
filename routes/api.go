package routes

import (
	"go-gin-file/app/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	{
		setUpRouter(router)
	}

	return router
}

func setUpRouter(router *gin.Engine) {
	a := router.Group("/api")
	{
		api.RegisterRouter(a)
	}
}
