package config

import (
	"FindMyDosen/model/entity"
	"fmt"
	"os"
)

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
