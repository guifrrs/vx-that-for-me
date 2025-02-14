package main

import (
	"log"

	"github.com/yanzay/tbot/v2"
)

func getUsername(msg *tbot.Message) string {
	if msg.From.Username != "" {
		return msg.From.Username
	}

	return msg.From.FirstName
}

func MessageHandler(msg *tbot.Message) {
	logMsg := "Received message from @" + getUsername(msg) + ": " + msg.Text
	log.Println(logMsg)

	tweetLink := "https://fixvx.com/" + msg.Text[20:]
	originalSenderMsg := "Hey @" + getUsername(msg) + ", I fixed that for you :3"

	app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	_, err := app.client.SendMessage(msg.Chat.ID, tweetLink)
	if err != nil {
		log.Println("Error sending message:", err)
	}

	app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
}
