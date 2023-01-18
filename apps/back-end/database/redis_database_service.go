package database

import (
	"FindMyDosen/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var client *redis.Client

func loadRedisDB() {
	url, err := getRedisURL()
	if err != nil {
		log.Fatal(err)
	}
	redisClient := redis.NewClient(url)
	client = redisClient
}

func getRedisURL() (*redis.Options, error) {
	conf := config.GetRedisConf()
	link := fmt.Sprintf("rediss://default:%s@%s:%s", conf.Password, conf.Address, conf.Port)
	return redis.ParseURL(link)
}

func GetRedisDB() *redis.Client {
	once.Do(loadRedisDB)
	return client
}
