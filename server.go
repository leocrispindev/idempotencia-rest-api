package main

import (
	"api-handle/internal/controller"
	"api-handle/internal/database"
	"api-handle/internal/services/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	InitConnections()

	r := gin.Default()

	controller.InitRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConnections() {
	database.Init()
	cache.Init()

}
