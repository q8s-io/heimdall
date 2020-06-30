package redis

import (
	"log"
)

func SetKV(key, value string) {
	_, keyErr := Client.Set(key, value, 0).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
}

func DelKV(key string) {
	_, keyErr := Client.Del(key).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
}

func SetMap(key, field string, value interface{}) {
	_, keyErr := Client.HSet(key, field, value).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
}

func GetMapAll(key string) map[string]string {
	data, keyErr := Client.HGetAll(key).Result()
	if keyErr != nil {
		log.Println(keyErr)
	}
	return data
}
