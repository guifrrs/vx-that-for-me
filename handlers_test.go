package main

import (
	"regexp"
	"testing"

	"github.com/yanzay/tbot/v2"
)

func TestGetUsername(t *testing.T) {
	testCases := []struct {
		name     string
		msg      *tbot.Message
		expected string
	}{
		{
			name: "Username exists",
			msg: &tbot.Message{
				From: &tbot.User{
					Username:  "johndoe",
					FirstName: "John",
				},
			},
			expected: "johndoe",
		},
		{
			name: "Username empty, use FirstName",
			msg: &tbot.Message{
				From: &tbot.User{
					Username:  "",
					FirstName: "John",
				},
			},
			expected: "John",
		},
		{
			name: "Username with special chars",
			msg: &tbot.Message{
				From: &tbot.User{
					Username:  "user_123-test",
					FirstName: "John",
				},
			},
			expected: "user_123-test",
		},
		{
			name:     "nil From field",
			msg:      &tbot.Message{},
			expected: "unknown",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getUsername(tc.msg)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestReplaceLink(t *testing.T) {
	testCases := []struct {
		name     string
		msg      *tbot.Message
		expected string
	}{
		{
			name: "Normal link",
			msg: &tbot.Message{
				Text: "https://x.com/whotfisjovana/status/1894867871412691219",
			},
			expected: "https://fixupx.com/whotfisjovana/status/1894867871412691219",
		},
		{
			name: "Twitter domain link",
			msg: &tbot.Message{
				Text: "https://twitter.com/theo/status/1895247223577026832",
			},
			expected: "https://fixupx.com/theo/status/1895247223577026832",
		},
		{
			name: "HTTP protocol link",
			msg: &tbot.Message{
				Text: "http://x.com/user/status/1895247223577026832",
			},
			expected: "https://fixupx.com/user/status/1895247223577026832",
		},
		{
			name: "WWW subdomain link",
			msg: &tbot.Message{
				Text: "https://www.x.com/user/status/1895247223577026832",
			},
			expected: "https://fixupx.com/user/status/1895247223577026832",
		},
		{
			name: "Link with query parameters",
			msg: &tbot.Message{
				Text: "https://x.com/user/status/1895247223577026832?s=20&t=abc",
			},
			expected: "https://fixupx.com/user/status/1895247223577026832?s=20&t=abc",
		},
		{
			name: "Link with anchor",
			msg: &tbot.Message{
				Text: "https://x.com/user/status/1895247223577026832#m",
			},
			expected: "https://fixupx.com/user/status/1895247223577026832#m",
		},
		{
			name: "Media page should not be converted",
			msg: &tbot.Message{
				Text: "https://x.com/DataFutebol/media",
			},
			expected: "https://x.com/DataFutebol/media",
		},
		{
			name: "Profile page should not be converted",
			msg: &tbot.Message{
				Text: "https://x.com/DataFutebol",
			},
			expected: "https://x.com/DataFutebol",
		},
		{
			name: "Text with multiple URLs - first is status",
			msg: &tbot.Message{
				Text: "Check this out https://x.com/user/status/123 and also https://x.com/other",
			},
			expected: "Check this out https://fixupx.com/user/status/123 and also https://x.com/other",
		},
		{
			name: "Text with multiple status URLs",
			msg: &tbot.Message{
				Text: "First: https://x.com/user1/status/123 Second: https://twitter.com/user2/status/456",
			},
			expected: "First: https://fixupx.com/user1/status/123 Second: https://fixupx.com/user2/status/456",
		},
		{
			name: "Instagram post link",
			msg: &tbot.Message{
				Text: "https://www.instagram.com/p/ABC123xyz/",
			},
			expected: "https://kkinstagram.com/p/ABC123xyz/",
		},
		{
			name: "Instagram reel link with query parameters",
			msg: &tbot.Message{
				Text: "Check this https://instagram.com/reel/ABC123/?igsh=example",
			},
			expected: "Check this https://kkinstagram.com/reel/ABC123/?igsh=example",
		},
		{
			name: "Instagram link followed by punctuation",
			msg: &tbot.Message{
				Text: "See https://instagram.com/p/ABC123.",
			},
			expected: "See https://kkinstagram.com/p/ABC123.",
		},
		{
			name: "Instagram TV link over HTTP",
			msg: &tbot.Message{
				Text: "http://instagram.com/tv/ABC123",
			},
			expected: "https://kkinstagram.com/tv/ABC123",
		},
		{
			name: "Instagram profile should not be converted",
			msg: &tbot.Message{
				Text: "https://instagram.com/example",
			},
			expected: "https://instagram.com/example",
		},
		{
			name: "Mixed Twitter and Instagram links",
			msg: &tbot.Message{
				Text: "https://x.com/user/status/123 https://instagram.com/p/ABC123/",
			},
			expected: "https://fixupx.com/user/status/123 https://kkinstagram.com/p/ABC123/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := replaceLink(tc.msg)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestSupportedLinkPattern(t *testing.T) {
	pattern := regexp.MustCompile(supportedLinkPattern())
	testCases := []struct {
		url       string
		supported bool
	}{
		{url: "https://x.com/user/status/123", supported: true},
		{url: "https://instagram.com/p/ABC123", supported: true},
		{url: "https://instagram.com/reel/ABC123", supported: true},
		{url: "https://instagram.com/example", supported: false},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			if actual := pattern.MatchString(tc.url); actual != tc.supported {
				t.Errorf("Expected supported=%t, got %t", tc.supported, actual)
			}
		})
	}
}

func TestSanitizeForLog(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal text",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "Text with newlines",
			input:    "Hello\nWorld\r\nTest",
			expected: "HelloWorldTest",
		},
		{
			name:     "Text with control characters",
			input:    "Hello\x00World\x01Test",
			expected: "HelloWorldTest",
		},
		{
			name:     "Text with special chars",
			input:    "user@example.com <script>",
			expected: "user@example.com script",
		},
		{
			name:     "Username with hyphen and underscore",
			input:    "user_name-test",
			expected: "user_name-test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sanitizeForLog(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
