package redis

import (
	"log"

	"github.com/pkg/errors"
)

func SetKV(key, value string) {
	keyStatus, keyErr := Client.Set(key, value, 0).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
	if keyStatus != "OK" {
		keyStatusErr := errors.New("write to redis faild")
		log.Println(keyStatusErr)
	}
}

func DelKV(key string) {
	_, keyErr := Client.Del(key).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
}

func SetMap(key, field string, value interface{}) {
	keyStatus, keyErr := Client.HSet(key, field, value).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
	if keyStatus == false {
		keyStatusErr := errors.New("write to redis faild")
		log.Println(keyStatusErr)
	}
}
