package controller

import (
	"api-handle/internal/model"
	"api-handle/internal/services/cache"
	"api-handle/internal/services/process"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var channelMessages chan model.Message

func Init() {
	process.Init()
	//channelMessages = process.Subscriber()
}

func GetList(c *gin.Context) {
	c.JSON(200, process.GetAll())

}

func RegistryMessage(c *gin.Context) {
	message := model.Message{}

	found, idempotenciaKey := hasIdempotenciakey(c.Request.Header)

	if !found {
		c.JSON(400, model.Error{
			Text: "Header idempotencia key not found",
		})
		return
	}

	err := c.BindJSON(&message.Info)

	if err != nil {
		c.JSON(400, model.Error{
			Text: "Error on parse body " + err.Error(),
		})
		return
	}

	message.IdempotenciaKey = idempotenciaKey

	cacheMessage, err := cache.IsOnCache(message.IdempotenciaKey)

	if err == nil && cacheMessage.InProccess() {
		c.JSON(200, nil)
	}

	err = process.ProccessMessage(message)

	if err != nil {
		c.JSON(500, model.Error{
			Text: "Registry message error",
		})
		return
	}

	c.JSON(204, nil)
}

func convertStringToUint(str string) (uint32, error) {

	result, err := strconv.ParseUint(str, 10, 32)

	if err != nil {
		return 0, err
	}
	return uint32(result), nil

}

func hasIdempotenciakey(headers http.Header) (bool, string) {
	result := headers.Get("Idempotencia-key")

	if result != "" {
		return true, result
	}

	return false, ""
}
