package main

import (
	"fmt"

	"log"

	"github.com/yanzay/tbot/v2"
)

func (app *Application) MessageHandler(msg *tbot.Message) {
	tweetLink := "https://fixvx.com/" + msg.Text[20:]

	var username string
	if msg.Chat.Username != "" {
		username = msg.From.Username
	} else {
		username = msg.From.FirstName
	}

	originalSenderMsg := fmt.Sprintf("Hey @%s, I fixed that for you :3", username)

	log.Println("Sender: ", msg.Chat.Username)

	app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	_, err := app.client.SendMessage(msg.Chat.ID, tweetLink)

	if err == nil {
		app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
	}
}
