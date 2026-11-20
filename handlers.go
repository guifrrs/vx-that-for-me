package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/yanzay/tbot/v2"
)

const (
	twitterLinkPattern   = `https?://(?:www\.)?(?:twitter|x)\.com(/[^/]+/status/\d+)`
	instagramLinkPattern = `https?://(?:www\.)?instagram\.com(/(?:p|reels?)/[A-Za-z0-9_-]+)`
	// Full post links and the /r/<sub>/s/<id> links the Reddit app shares.
	redditLinkPattern      = `https?://(?:(?:www|old|new|np|m)\.)?reddit\.com(/(?:(?:(?:r|u|user)/[^/\s]+/)?comments/[a-z0-9]+|(?:r|u|user)/[^/\s]+/s/[A-Za-z0-9]+))`
	redditShortLinkPattern = `https?://redd\.it(/[a-z0-9]+)`
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
	newLinkRule(instagramLinkPattern, "zzinstagram.com"),
	newLinkRule(redditLinkPattern, "vxreddit.com"),
	newLinkRule(redditShortLinkPattern, "vxreddit.com"),
}

// Telegram keeps unconfirmed updates for up to 24h, so after the bot was offline
// it receives everything posted meanwhile. Links older than this are left alone.
const maxMessageAge = 2 * time.Minute

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

func isStale(msg *tbot.Message, now time.Time) bool {
	return now.Sub(time.Unix(msg.Date, 0)) > maxMessageAge
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

	if isStale(msg, time.Now()) {
		log.Printf("Skipping stale message from @%s sent at %s", safeUsername, time.Unix(msg.Date, 0).Format(time.RFC3339))
		return
	}

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
