package cache

import (
	"api-handle/internal/model"
	"api-handle/internal/services/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Fail on connect to Redis: %v", err)
	}

	log.Println("Cache ok")
}

func Set(key string, value string, expiration int) error {
	if utils.IsStringEmpty(value) || utils.IsStringEmpty(key) || expiration < 0 {
		return fmt.Errorf("Key or value is empty")
	}

	durationUtilExpire := time.Duration(expiration) * time.Second

	return rdb.Set(context.TODO(), key, value, durationUtilExpire).Err()

}

func SetNotExist(key string, value string, expiration int) error {
	if utils.IsStringEmpty(value) || utils.IsStringEmpty(key) || expiration < 0 {
		return fmt.Errorf("Key or value is empty")
	}

	durationUtilExpire := time.Duration(expiration) * time.Second

	return rdb.SetNX(context.TODO(), key, value, durationUtilExpire).Err()

}

func Retrieve(key string) (string, error) {
	val, err := rdb.Get(context.TODO(), key).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")

	} else if err != nil {
		return "", fmt.Errorf("Error on retrieve value from redis: %v", err)

	}

	return val, nil
}

func IsOnCache(idempotenciaKey string) (model.CacheMessage, error) {
	var result model.CacheMessage

	data, err := Retrieve(idempotenciaKey)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(data), &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
