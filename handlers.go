package main

import (
	"fmt"

	"github.com/yanzay/tbot/v2"
)

func getUsername(msg *tbot.Message) string {
	if msg.Chat.Username != "" {
		return msg.From.Username
	}

	if msg.From.LastName != "" {
		return msg.From.FirstName + " " + msg.From.LastName
	}

	return msg.From.FirstName
}

func (app *Application) MessageHandler(msg *tbot.Message) {
	tweetLink := "https://fixvx.com/" + msg.Text[20:]
	originalSenderMsg := fmt.Sprintf("Hey @%s, I fixed that for you :3", getUsername(msg))

	app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	_, err := app.client.SendMessage(msg.Chat.ID, tweetLink)

	if err == nil {
		app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
	}
}
