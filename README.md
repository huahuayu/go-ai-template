# Go AI Template

A simple, dependency-free Go template for calling LLM models (Gemini & Claude) via an Antigravity-compatible API gateway.

## Features

- **Protocol Native**: Calls Gemini via Google REST and Claude via Anthropic Messages format.
- **Dependency Free**: Uses only Go standard library (`net/http`, `encoding/json`).
- **Antigravity Optimized**: Works perfectly with `sub2api` or direct Antigravity endpoints.

## Models Covered

- **Gemini**: `gemini-3-flash`, `gemini-3-pro-low`, `gemini-3-pro-high`, `gemini-2.5-pro`, etc.
- **Claude**: `claude-3-5-sonnet-20240620`, `claude-sonnet-4-5`, `claude-opus-4-5-thinking` (High-end model), etc.

## Usage

1. Set your environment variables:
   ```bash
   export AI_BASE_URL="https://model.liushiming.cn/antigravity"
   export AI_API_KEY="your-sk-key"
   ```

2. Run the code:
   ```bash
   go run main.go
   ```

## Key Concept: Dual Protocol Support

The template demonstrates how to handle different API schemas:
- **Gemini Path**: `/v1beta/models/{model}:generateContent`
- **Claude Path**: `/v1/messages`

By using the `/antigravity` path in your Base URL, the gateway correctly routes both protocols using a single API Key.
