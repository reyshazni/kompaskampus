package database

import (
	"FindMyDosen/config"
	"FindMyDosen/model/entity"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB
var once sync.Once

func loadDB() {
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: getDatabaseDsn(),
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	//gorm.Open("mysql", getDatabaseDsn())
	if err != nil {
		log.Fatal("Error loading Database", err)
	}
	// defer db.Close()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Lecture{})
	db.AutoMigrate(&entity.University{})
	db.AutoMigrate(&entity.Subject{})
	db.AutoMigrate(&entity.LectureSubject{})
	db.AutoMigrate(&entity.LectureRating{})
	db.AutoMigrate(&entity.RefreshToken{})
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

func Paginate(page int, limit int) *gorm.DB {
	once.Do(loadDB)
	pageLimit := limit
	currPage := page
	if currPage <= 0 {
		currPage = 1
	}
	switch {
	case pageLimit > 100:
		pageLimit = 100
	case pageLimit <= 0:
		pageLimit = 1
	}
	offset := (currPage - 1) * pageLimit
	return database.Offset(offset).Limit(pageLimit)
}
