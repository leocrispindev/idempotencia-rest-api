package process

import (
	messageDao "api-handle/internal/dao/message"
	"api-handle/internal/model"
	"api-handle/internal/services/cache"
	"log"
)

func GetAll() []interface{} {
	return messageDao.GetAllMessages()
}

func ProccessMessage(message model.Message) error {
	if messageAlreadyProcessed(message.IdempotenciaKey) {
		return nil
	}

	saveOnCache(message.IdempotenciaKey, model.IN_PROCESS, "Message in process", 6) // seconds

	err := messageDao.SaveMessage(message)

	if err != nil {
		log.Printf("Error on save message, [ID]= %s  [error]=%s", message.IdempotenciaKey, err.Error())

		saveOnCache(message.IdempotenciaKey, model.ERROR_ON_PROCESS, "Error on save message", 60)

	}

	saveOnCache(message.IdempotenciaKey, model.PROCESSED, "Save with success", 60)

	return err
}

func messageAlreadyProcessed(key string) bool {

	result, err := messageDao.AlreadyProcessed(key)

	if err != nil {
		log.Println(err.Error())
	}

	return result

}

func saveOnCache(key string, status model.Status, message string, duration int) {
	cacheMessage := model.CacheMessage{
		ProcessStatus: status,
		Message:       message,
	}

	cache.Set(key, cacheMessage.ToJson(), duration*1000)

}
