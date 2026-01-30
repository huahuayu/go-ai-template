package main

import (
	"encoding/json"
	"fmt"
)

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

	return c.DoRequest(url, reqBody, func(body []byte) (string, error) {
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
