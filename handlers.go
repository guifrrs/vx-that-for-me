package main

import "github.com/yanzay/tbot/v2"

func (app *Application) MessageHandler(msg *tbot.Message) {
	tweetLink := "https://fixvx.com/" + msg.Text[20:]

	app.client.SendMessage(msg.Chat.ID, "Here, I fixed that for you :3")
	_, err := app.client.SendMessage(msg.Chat.ID, tweetLink)

	if err == nil {
		app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
	}
}
