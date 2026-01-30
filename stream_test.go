package main

import (
	"os"
	"testing"
)

func TestStreamGemini(t *testing.T) {
	baseURL := os.Getenv("AI_BASE_URL")
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		t.Skip("Skipping test: AI_BASE_URL or AI_API_KEY not set")
	}

	client := NewClient(baseURL, apiKey)

	// 测试 Gemini 3 Flash 的流式输出
	err := client.StreamGemini("gemini-3-flash", "Tell me a very short joke.")
	if err != nil {
		t.Errorf("Stream Error: %v", err)
	}
}
