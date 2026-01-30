package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Config holds the API configuration
type Config struct {
	BaseURL string
	APIKey  string
}

// Message represents a single chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// --- Gemini (Google SDK Format) Structures ---

type GeminiContent struct {
	Role  string       `json:"role,omitempty"`
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// --- Claude (Anthropic Format) Structures ---

type ClaudeRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type ClaudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

// Client represents our AI service client
type Client struct {
	cfg Config
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{cfg: Config{BaseURL: baseURL, APIKey: apiKey}}
}

// CallGemini calls a Gemini model using the Google REST protocol
func (c *Client) CallGemini(modelName, prompt string) (string, error) {
	url := fmt.Sprintf("%s/v1beta/models/%s:generateContent", c.cfg.BaseURL, modelName)

	reqBody := GeminiRequest{
		Contents: []GeminiContent{
			{
				Role:  "user",
				Parts: []GeminiPart{{Text: prompt}},
			},
		},
	}

	return c.doRequest(url, reqBody, func(body []byte) (string, error) {
		var resp GeminiResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			return "", err
		}
		if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			return resp.Candidates[0].Content.Parts[0].Text, nil
		}
		return "", fmt.Errorf("empty gemini response")
	})
}

// CallClaude calls a Claude model using the Anthropic Messages protocol
func (c *Client) CallClaude(modelName, prompt string) (string, error) {
	url := fmt.Sprintf("%s/v1/messages", c.cfg.BaseURL)

	reqBody := ClaudeRequest{
		Model: modelName,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024,
	}

	return c.doRequest(url, reqBody, func(body []byte) (string, error) {
		var resp ClaudeResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			return "", err
		}
		if len(resp.Content) > 0 {
			return resp.Content[0].Text, nil
		}
		return "", fmt.Errorf("empty claude response")
	})
}

func (c *Client) doRequest(url string, payload interface{}, parser func([]byte) (string, error)) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.APIKey)

	// Note: Antigravity gateway supports both protocols via the /antigravity path
	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	return parser(body)
}

func main() {
	// 1. Setup Configuration
	// Recommended: use environment variables for keys
	baseURL := os.Getenv("AI_BASE_URL") // e.g., https://model.liushiming.cn/antigravity
	apiKey := os.Getenv("AI_API_KEY")

	if baseURL == "" || apiKey == "" {
		fmt.Println("Usage: AI_BASE_URL=... AI_API_KEY=... go run main.go")
		return
	}

	client := NewClient(baseURL, apiKey)

	// 2. Test Gemini 3 Pro High
	fmt.Println("--- Testing Gemini 3 Pro High ---")
	geminiText, err := client.CallGemini("gemini-3-pro-high", "What is the capital of France?")
	if err != nil {
		fmt.Printf("Gemini Error: %v\n", err)
	} else {
		fmt.Printf("Gemini Response: %s\n", geminiText)
	}

	fmt.Println()

	// 3. Test Claude 3.5 Sonnet
	fmt.Println("--- Testing Claude 3.5 Sonnet ---")
	claudeText, err := client.CallClaude("claude-3-5-sonnet-20240620", "Write a short poem about coding.")
	if err != nil {
		fmt.Printf("Claude Error: %v\n", err)
	} else {
		fmt.Printf("Claude Response: %s\n", claudeText)
	}

	fmt.Println()

	// 4. Test Claude 4.5 Opus
	fmt.Println("--- Testing Claude 4.5 Opus ---")
	opusText, err := client.CallClaude("claude-opus-4-5-thinking", "Explain the concept of concurrency in Go in one sentence.")
	if err != nil {
		fmt.Printf("Opus Error: %v\n", err)
	} else {
		fmt.Printf("Opus Response: %s\n", opusText)
	}
}
