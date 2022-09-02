package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sergot/tibiago/src/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	fmt.Println("Hello, World!")

	bot.Connect()
}
