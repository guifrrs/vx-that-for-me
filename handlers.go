package main

import "github.com/yanzay/tbot/v2"

func (app *Application) MessageHandler(msg *tbot.Message) {
	tweetLink := "https://fixvx.com/" + msg.Text[20:]

	app.client.SendMessage(msg.Chat.ID, "Here, I fixed that for you")
	app.client.SendMessage(msg.Chat.ID, tweetLink)
}
