package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/huahuayu/go-ai-template/ai"
)

func TestAIModels(t *testing.T) {
	baseURL := os.Getenv("AI_BASE_URL")
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		t.Skip("Skipping test: AI_BASE_URL or AI_API_KEY not set")
	}

	client := ai.NewClient(baseURL, apiKey)

	t.Run("Gemini-3-Flash", func(t *testing.T) {
		resp, err := client.CallGemini("gemini-3-flash", "Hello")
		if err != nil { t.Errorf("Error: %v", err) }
		fmt.Printf("[Gemini 3 Flash] %s\n", resp)
	})

	t.Run("Claude-Sonnet-4.5", func(t *testing.T) {
		resp, err := client.CallClaude("claude-sonnet-4-5", "Hello")
		if err != nil { t.Errorf("Error: %v", err) }
		fmt.Printf("[Claude Sonnet 4.5] %s\n", resp)
	})

	t.Run("Claude-Opus-4.5-Standard", func(t *testing.T) {
		resp, err := client.CallClaude("claude-opus-4-5", "Hi")
		if err != nil { t.Errorf("Error: %v", err) }
		fmt.Printf("[Claude Opus 4.5] %s\n", resp)
	})

	t.Run("Claude-Opus-4.5-Thinking", func(t *testing.T) {
		resp, err := client.CallClaude("claude-opus-4-5-thinking", "Explain quantum in 5 words")
		if err != nil { t.Errorf("Error: %v", err) }
		fmt.Printf("[Claude Opus 4.5 Thinking] %s\n", resp)
	})
}
