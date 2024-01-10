package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccessResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"text": message,
	})
	c.Abort()
}

func sendErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
	c.Abort()
}

func sendBadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
	c.Abort()
}
