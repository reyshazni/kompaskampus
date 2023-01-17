package main

import (
	"FindMyDosen/application"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func Main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	application.ApplicationDelegate()
}

func main() {
	Main()
}
