package process

import (
	messageDao "api-handle/internal/dao/message"
	"api-handle/internal/model"
	"api-handle/internal/services/cache"
	"log"
)

func Init() {

}

func GetAll() []interface{} {
	return messageDao.GetAllMessages()
}

func ProccessMessage(message model.Message) error {
	if messageAlreadyProcessed(message.IdempotenciaKey) {
		return nil
	}

	saveOnCache(message.IdempotenciaKey, model.IN_PROCESS, "Message in process", 600) // 10 minutes cache duration

	err := messageDao.SaveMessage(message)

	if err != nil {
		log.Printf("Error on save message, [ID]= %s  [error]=%s", message.IdempotenciaKey, err.Error())

		go saveOnCache(message.IdempotenciaKey, model.ERROR_ON_PROCESS, "Error on save message", 600) // 10 minutes cache duration

	}

	go saveOnCache(message.IdempotenciaKey, model.PROCESSED, "Prcessed with success", 60)

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

	go cache.Set(key, cacheMessage.ToJson(), duration)
}
