package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGeminiModels(t *testing.T) {
	baseURL := os.Getenv("AI_BASE_URL")
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		t.Skip("Skipping test: AI_BASE_URL or AI_API_KEY not set")
	}

	client := NewClient(baseURL, apiKey)

	models := []string{
		"gemini-3-flash",
		"gemini-3-pro-low",
		"gemini-3-pro-high",
		"gemini-2.5-flash",
	}

	for _, m := range models {
		t.Run(m, func(t *testing.T) {
			resp, err := client.CallGemini(m, "Say 'Gemini is online'")
			if err != nil {
				t.Errorf("Error calling %s: %v", m, err)
			} else {
				fmt.Printf("[%s] Response: %s\n", m, resp)
			}
		})
	}
}
