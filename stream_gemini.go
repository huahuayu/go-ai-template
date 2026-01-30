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

// GeminiStreamResponse represents the Google Gemini SSE chunk format
type GeminiStreamResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func (c *Client) StreamGemini(modelName, prompt string) error {
	// Gemini uses streamGenerateContent for streaming
	url := fmt.Sprintf("%s/v1beta/models/%s:streamGenerateContent?alt=sse", c.cfg.BaseURL, modelName)

	reqBody := map[string]interface{}{
		"contents": []interface{}{
			map[string]interface{}{
				"role": "user",
				"parts": []interface{}{
					map[string]string{"text": prompt},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Gemini API error (%d): %s", resp.StatusCode, string(body))
	}

	// Use a scanner to read the SSE data line by line
	scanner := bufio.NewScanner(resp.Body)
	fmt.Printf("[%s] Streaming: ", modelName)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			break
		}

		var streamResp GeminiStreamResponse
		if err := json.Unmarshal([]byte(data), &streamResp); err != nil {
			continue
		}

		if len(streamResp.Candidates) > 0 && len(streamResp.Candidates[0].Content.Parts) > 0 {
			content := streamResp.Candidates[0].Content.Parts[0].Text
			fmt.Print(content) // Print each chunk as it arrives
		}
	}

	fmt.Println("\n[Gemini Stream Finished]")
	return scanner.Err()
}
