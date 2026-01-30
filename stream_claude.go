package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ClaudeStreamDelta represents the data in a content_block_delta event
type ClaudeStreamDelta struct {
	Type string `json:"type"`
	Delta struct {
		Text string `json:"text"`
	} `json:"delta"`
}

func (c *Client) StreamClaude(modelName, prompt string) error {
	url := fmt.Sprintf("%s/v1/messages", c.cfg.BaseURL)

	reqBody := map[string]interface{}{
		"model":      modelName,
		"messages":   []map[string]string{{"role": "user", "content": prompt}},
		"max_tokens": 1024,
		"stream":     true,
	}

	jsonData, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Claude API error (%d): %s", resp.StatusCode, string(body))
	}

	scanner := bufio.NewScanner(resp.Body)
	fmt.Printf("[%s] Streaming: ", modelName)
	for scanner.Scan() {
		line := scanner.Text()
		
		// Claude uses "event: ..." followed by "data: ..."
		// We only care about data lines in this simplified example
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		data := strings.TrimPrefix(line, "data: ")
		
		var delta ClaudeStreamDelta
		if err := json.Unmarshal([]byte(data), &delta); err != nil {
			continue
		}

		// Only print if it's a content_block_delta
		if delta.Type == "content_block_delta" {
			fmt.Print(delta.Delta.Text)
		}
	}

	fmt.Println("\n[Claude Stream Finished]")
	return scanner.Err()
}
