package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	SetupMessagesRoutes(r)
}
