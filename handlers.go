package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/yanzay/tbot/v2"
)

const (
	twitterLinkPattern   = `https?://(?:www\.)?(?:twitter|x)\.com(/[^/]+/status/\d+)`
	instagramLinkPattern = `https?://(?:www\.)?instagram\.com(/(?:p|reels?|tv)/[A-Za-z0-9_-]+)`
)

type linkRule struct {
	pattern     string
	regex       *regexp.Regexp
	replacement string
}

func newLinkRule(pattern, domain string) linkRule {
	return linkRule{
		pattern:     pattern,
		regex:       regexp.MustCompile(pattern),
		replacement: "https://" + domain + "$1",
	}
}

var linkRules = []linkRule{
	newLinkRule(twitterLinkPattern, "fixupx.com"),
	newLinkRule(instagramLinkPattern, "kkinstagram.com"),
}

var usernameRegex = regexp.MustCompile(`[^a-zA-Z0-9_\-@. ]`)

func supportedLinkPattern() string {
	patterns := make([]string, 0, len(linkRules))
	for _, rule := range linkRules {
		patterns = append(patterns, rule.pattern)
	}
	return `(?:` + strings.Join(patterns, `|`) + `)`
}

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
	result := msg.Text
	for _, rule := range linkRules {
		result = rule.regex.ReplaceAllString(result, rule.replacement)
	}
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
