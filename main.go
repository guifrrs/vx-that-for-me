package main

import (
	"fmt"
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

func getUsername(msg *tbot.Message) string {
	if msg.From.Username != "" {
		return msg.From.Username
	}

	return msg.From.FirstName
}

func MessageHandler(msg *tbot.Message) {
	tweetLink := "https://fixvx.com/" + msg.Text[20:]
	originalSenderMsg := fmt.Sprintf("Hey @%s, I fixed that for you :3", getUsername(msg))

	app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	_, err := app.client.SendMessage(msg.Chat.ID, tweetLink)

	if err == nil {
		app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
	}
}

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
