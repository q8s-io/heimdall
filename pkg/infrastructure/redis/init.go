package redis

import (
	"time"

	"github.com/go-redis/redis"

	"github.com/q8s-io/heimdall/pkg/models"
)

var Client *redis.Client

func Init() {
	redisConfig := models.Config.Redis
	Client = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Host + ":" + redisConfig.Port,
		Password:     redisConfig.PassWord,
		MaxRetries:   2,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  5 * time.Second,
		DB:           0,
		PoolSize:     5,
		IdleTimeout:  15 * time.Second,
	})

	go func() {
		taskConnect := time.NewTicker(3 * time.Second)
		for {
			<-taskConnect.C
			go Status()
		}
	}()
}
