package main

import (
	"FindMyDosen/database"
	"FindMyDosen/model/entity"
	"github.com/go-faker/faker/v4"
	"log"
)

func fakeLecture() {
	lecture := []entity.Lecture{}
	db := database.GetDB()
	if err := faker.FakeData(&lecture); err != nil {
		log.Fatal(err)
	}
	if err := db.Create(&lecture).Error; err != nil {
		log.Fatal(err)
	}
}
