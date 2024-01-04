package main

import (
	"api-handle/internal/controller"
	"api-handle/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	controller.Init()

	r := gin.Default()

	r.GET("/rest/list/all", controller.GetList)

	// Without Idempotencia
	r.POST("/rest/registry", controller.RegistryMessage)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
