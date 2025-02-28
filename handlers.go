package main

import (
	"log"
	"regexp"

	"github.com/yanzay/tbot/v2"
)

func getUsername(msg *tbot.Message) string {
	if msg.From.Username != "" {
		return msg.From.Username
	}

	return msg.From.FirstName
}

func replaceLink(msg *tbot.Message) string {
	domainRegex := regexp.MustCompile(`https://(.*?)/`)
	domain := domainRegex.FindString(msg.Text)
	if domain != "" {
		return "https://fixvx.com/" + msg.Text[len(domain):]
	}

	return msg.Text
}

func MessageHandler(msg *tbot.Message) {
	logMsg := "Received message from @" + getUsername(msg) + ": " + msg.Text
	log.Println(logMsg)

	originalSenderMsg := "Hey @" + getUsername(msg) + ", I fixed that for you :3"
	app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	_, err := app.client.SendMessage(msg.Chat.ID, replaceLink(msg))
	if err != nil {
		log.Println("Error sending message:", err)
	}

	app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
}
