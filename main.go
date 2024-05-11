package main

import (
	"go-gin-file/config"
	"go-gin-file/routes"
)

func main() {
	port := config.GetInitKeyFatal("server", "SERVER_PORT")

	r := routes.InitRouter()
	r.Run(":" + port)
}
