package controller

import (
	"api-handle/internal/model"
	"api-handle/internal/services/cache"
	"api-handle/internal/services/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	process.Init()
}

func GetList(c *gin.Context) {
	c.JSON(200, process.GetAll())

}

func RegistryMessage(c *gin.Context) {
	message := model.Message{}

	found, idempotenciaKey := hasIdempotenciakey(c.Request.Header)

	if !found {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Text: "Header idempotencia key not found",
		})
		return
	}

	err := c.BindJSON(&message.Info)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Text: "Error on parse body " + err.Error(),
		})
		return
	}

	message.IdempotenciaKey = idempotenciaKey

	cacheMessage, err := cache.IsOnCache(message.IdempotenciaKey)

	if err == nil && cacheMessage.InProccess() {
		c.JSON(http.StatusOK, model.ResponseSuccess{
			Text: cacheMessage.Message,
		})
		return
	}

	if cacheMessage.StatusError() {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Text: cacheMessage.Message,
		})
		return
	}

	err = process.ProccessMessage(message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Text: "Registry message error",
		})
		return
	}

	c.JSON(http.StatusCreated, c)
}

func hasIdempotenciakey(headers http.Header) (bool, string) {
	result := headers.Get("Idempotencia-key")

	if result != "" {
		return true, result
	}

	return false, ""
}
