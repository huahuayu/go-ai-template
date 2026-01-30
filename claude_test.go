package main

import (
	"fmt"
	"os"
	"testing"
)

func TestClaudeModels(t *testing.T) {
	baseURL := os.Getenv("AI_BASE_URL")
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		t.Skip("Skipping test: AI_BASE_URL or AI_API_KEY not set")
	}

	client := NewClient(baseURL, apiKey)

	models := []string{
		"claude-3-5-sonnet-20240620",
		"claude-sonnet-4-5",
		"claude-opus-4-5",
		"claude-opus-4-5-thinking",
	}

	for _, m := range models {
		t.Run(m, func(t *testing.T) {
			resp, err := client.CallClaude(m, "Say 'Claude is online'")
			if err != nil {
				t.Errorf("Error calling %s: %v", m, err)
			} else {
				fmt.Printf("[%s] Response: %s\n", m, resp)
			}
		})
	}
}
