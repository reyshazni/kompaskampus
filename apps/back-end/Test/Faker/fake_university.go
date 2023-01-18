package main

import (
	"FindMyDosen/database"
	"FindMyDosen/model/entity"
	"github.com/go-faker/faker/v4"
	"log"
)

func fakeUniversity() {
	uni := []entity.University{}
	db := database.GetDB()
	if err := faker.FakeData(&uni); err != nil {
		log.Fatal(err)
	}
	if err := db.Create(&uni).Error; err != nil {
		log.Fatal(err)
	}
}
