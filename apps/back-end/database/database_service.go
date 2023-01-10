package database

import (
	"FindMyDosen/config"
	"FindMyDosen/model/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var database *gorm.DB
var once sync.Once

func loadDB() {
	db, err := gorm.Open("mysql", getDatabaseDsn())
	if err != nil {
		log.Fatal("Error loading Database", err)
	}
	// defer db.Close()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Lecture{})
	db.AutoMigrate(&entity.University{})
	database = db
}

func GetDB() *gorm.DB {
	once.Do(loadDB)
	return database
}

func getDatabaseDsn() string {
	entity := config.GetDatabaseConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", entity.User, entity.Password, entity.Url, entity.Port, entity.DbName)
	return dsn
}
