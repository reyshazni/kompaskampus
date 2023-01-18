package config

import (
	"FindMyDosen/model/entity"
	"fmt"
	"os"
)

type RedisEntity struct {
	Address  string
	Port     string
	Password string
}

func GetRedisConf() RedisEntity {
	return RedisEntity{
		Address:  os.Getenv("REDIS_ADR"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}

func GetDatabaseConf() entity.DatabaseEntity {
	return entity.DatabaseEntity{
		User:     os.Getenv("DATABASE_UNAME"),
		Password: os.Getenv("DATABASE_PWD"),
		DbName:   os.Getenv("DATABASE_DBNAME"),
		Url:      os.Getenv("DATABASE_URL"),
		Dialect:  os.Getenv("DATABASE_DIALECT"),
		Port:     os.Getenv("DATABASE_PORT"),
	}
}

func GetServerPort() string {
	port := os.Getenv("SERVER_PORT")
	return fmt.Sprintf(":%s", port)
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
