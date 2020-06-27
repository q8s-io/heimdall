package redis

import (
	"log"
)

func Status() {
	pong, err := Client.Ping().Result()
	if err != nil || pong != "PONG" {
		log.Println(err)
	}
}
