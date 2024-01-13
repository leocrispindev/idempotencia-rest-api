package process

import (
	messageDao "api-handle/internal/dao/message"
	"api-handle/internal/model"
	"api-handle/internal/services/cache"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func GetAll() []interface{} {
	return messageDao.GetAllMessages()
}

func ProccessMessage(message model.Message) error {
	cacheMessage, err := cache.IsOnCache(message.IdempotenciaKey)

	if err == nil && cacheMessage.StatusError() {
		return fmt.Errorf(cacheMessage.Message)
	}

	if (err == nil && cacheMessage.InProccess()) || messageAlreadyProcessed(message.IdempotenciaKey) {
		return nil
	}

	cacheMessage = createCacheMessageModel(model.IN_PROCESS, "Message in process")

	err = cache.SetNotExist("inproccess:"+message.IdempotenciaKey, cacheMessage.ToJson(), 6*1000) // seconds

	//Key already exist on redis
	if err == redis.Nil {
		return nil
	}

	err = messageDao.SaveMessage(message)

	if err != nil {
		log.Printf("Error on save message, [ID]= %s  [error]=%s", message.IdempotenciaKey, err.Error())

		cacheMessage.UpdateStatus(model.ERROR_ON_PROCESS, "Error on save message")

		cache.Set(message.IdempotenciaKey, cacheMessage.ToJson(), 60*1000)

	}

	cacheMessage.UpdateStatus(model.PROCESSED, "Save with success")
	cache.Set(message.IdempotenciaKey, cacheMessage.ToJson(), 60*1000)

	return err
}

func messageAlreadyProcessed(key string) bool {

	result, err := messageDao.AlreadyProcessed(key)

	if err != nil {
		log.Println(err.Error())
	}

	return result

}

func createCacheMessageModel(status model.Status, message string) model.CacheMessage {
	return model.CacheMessage{
		ProcessStatus: status,
		Message:       message,
	}
}
