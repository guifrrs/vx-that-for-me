package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type Application struct {
	client *tbot.Client
}

var (
	app   Application
	token string
	bot   *tbot.Server
)

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Println("Error loading .env file")
	}

	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	log.Println("Starting bot...")

	bot = tbot.New(token)
	app.client = bot.Client()

	bot.HandleMessage(`^https?:\/\/(?:www\.)?(?:twitter|x)\.com\/.*$`, MessageHandler)
	log.Fatal(bot.Start())
}
