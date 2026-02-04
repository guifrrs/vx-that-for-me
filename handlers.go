package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/yanzay/tbot/v2"
)

const fixupxDomain = "fixupx.com"

var (
	statusRegex   = regexp.MustCompile(`https?://(?:www\.)?(?:twitter|x)\.com/[^/]+/status/\d+`)
	domainRegex   = regexp.MustCompile(`https?://(.*?)/`)
	usernameRegex = regexp.MustCompile(`[^a-zA-Z0-9_\-@. ]`)
)

func sanitizeForLog(input string) string {
	return usernameRegex.ReplaceAllString(input, "")
}

func getUsername(msg *tbot.Message) string {
	if msg.From == nil {
		return "unknown"
	}
	if msg.From.Username != "" {
		return msg.From.Username
	}
	return msg.From.FirstName
}

func replaceLink(msg *tbot.Message) string {
	if !statusRegex.MatchString(msg.Text) {
		return msg.Text
	}

	result := statusRegex.ReplaceAllStringFunc(msg.Text, func(match string) string {
		domainMatch := domainRegex.FindStringSubmatch(match)
		if len(domainMatch) > 1 {
			return fmt.Sprintf("https://%s%s", fixupxDomain, match[len(domainMatch[0])-1:])
		}
		return match
	})

	return result
}

func MessageHandler(msg *tbot.Message) {
	username := getUsername(msg)
	safeUsername := sanitizeForLog(username)
	safeText := sanitizeForLog(msg.Text)
	log.Printf("Received message from @%s: %s", safeUsername, safeText)

	originalSenderMsg := fmt.Sprintf("Hey @%s, I fixed that for you :3", username)
	_, err := app.client.SendMessage(msg.Chat.ID, originalSenderMsg)
	if err != nil {
		log.Printf("Error sending notification message: %v", err)
		return
	}

	fixedLink := replaceLink(msg)
	_, err = app.client.SendMessage(msg.Chat.ID, fixedLink)
	if err != nil {
		log.Printf("Error sending fixed link: %v", err)
		return
	}

	err = app.client.DeleteMessage(msg.Chat.ID, msg.MessageID)
	if err != nil {
		log.Printf("Error deleting original message: %v", err)
	}
}
