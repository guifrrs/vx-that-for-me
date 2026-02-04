package main

import (
	"testing"

	"github.com/yanzay/tbot/v2"
)

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
			name: "Broken link",
			msg: &tbot.Message{
				Text: "https://x.com/theo/status/1895247223577026832",
			},
			expected: "https://fixupx.com/theo/status/1895247223577026832",
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
