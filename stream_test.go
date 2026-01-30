package main

import (
	"os"
	"testing"
)

func TestStreaming(t *testing.T) {
	baseURL := os.Getenv("AI_BASE_URL")
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		t.Skip("Skipping test: AI_BASE_URL or AI_API_KEY not set")
	}

	client := NewClient(baseURL, apiKey)

	t.Run("Gemini-3-Flash", func(t *testing.T) {
		err := client.StreamGemini("gemini-3-flash", "Say 'Gemini Stream is active'")
		if err != nil {
			t.Errorf("Gemini Stream Error: %v", err)
		}
	})

	t.Run("Claude-3-5-Sonnet", func(t *testing.T) {
		err := client.StreamClaude("claude-3-5-sonnet-20240620", "Say 'Claude Stream is active'")
		if err != nil {
			t.Errorf("Claude Stream Error: %v", err)
		}
	})
}
