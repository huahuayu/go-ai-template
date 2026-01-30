package ai

import (
	"encoding/json"
	"fmt"
)

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

func (c *Client) CallClaude(modelName, prompt string) (string, error) {
	url := fmt.Sprintf("%s/v1/messages", c.cfg.BaseURL)

	reqBody := ClaudeRequest{
		Model: modelName,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1024,
	}

	return c.DoRequest(url, reqBody, func(body []byte) (string, error) {
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
