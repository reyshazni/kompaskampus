package main

import (
	"FindMyDosen/application"
	"log"

	"github.com/joho/godotenv"
)

func Main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	application.ApplicationDelegate()
}

func main() {
	Main()
}
