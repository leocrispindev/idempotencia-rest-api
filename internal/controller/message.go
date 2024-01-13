package controller

import (
	"api-handle/internal/model"
	"api-handle/internal/services/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupMessagesRoutes(r *gin.Engine) {
	r.GET("/rest/list/all", getList)
	r.POST("/rest/registry", registryMessage)
}

func getList(c *gin.Context) {
	c.JSON(200, process.GetAll())
}

func registryMessage(c *gin.Context) {

	found, idempotenciaKey := hasIdempotenciakey(c.Request.Header)

	if !found {
		sendBadRequestResponse(c, "Key not found")
	}

	message := model.Message{}
	message.IdempotenciaKey = idempotenciaKey

	err := c.BindJSON(&message.Info)

	if err != nil {
		sendBadRequestResponse(c, "Error on parse body "+err.Error())
	}

	err = process.ProccessMessage(message)

	if err != nil {
		sendErrorResponse(c, "Registry message error")
	}

	sendSuccessResponse(c, "Success")
	c.Abort()
}

func hasIdempotenciakey(headers http.Header) (bool, string) {
	result := headers.Get("Idempotencia-key")

	if result != "" {
		return true, result
	}

	return false, ""
}
