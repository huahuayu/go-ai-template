# Go AI Template

A simple, modular Go template for calling LLM models (Gemini & Claude) via an Antigravity-compatible API gateway.

## Features

- **Modular Design**: Separated packages for `gemini` and `claude` logic.
- **Dependency Free**: Uses only Go standard library.
- **Test-Driven**: Examples provided as Go Tests for easy integration.

## Supported Models

### Claude
- `claude-opus-4-5-thinking` (Thinking Mode)
- `claude-opus-4-5` (Standard Mode)
- `claude-sonnet-4-5`
- `claude-sonnet-4-5-thinking`

### Gemini
- `gemini-3-flash`
- `gemini-3-pro-low`
- `gemini-3-pro-high`
- `gemini-3-pro-preview`
- `gemini-3-pro-image`
- `gemini-2.5-flash`
- `gemini-2.5-flash-lite`
- `gemini-2.5-flash-thinking`
- `gemini-2.5-pro`

## How to Get All Supported Models

You can fetch the real-time list of all models supported by your Antigravity accounts by calling the following endpoint:

```bash
curl -H "Authorization: Bearer your-sk-key" https://model.liushiming.cn/antigravity/models
```

This returns a JSON list of all available model IDs, which can be used in your API requests.

## Usage

1. Set your environment variables:
   ```bash
   export AI_BASE_URL="https://model.liushiming.cn/antigravity"
   export AI_API_KEY="your-sk-key"
   ```

2. Run the tests:
   ```bash
   go test -v .
   ```

## Key Concept: Dual Protocol Support

The template demonstrates how to handle different API schemas:
- **Gemini Path**: `/v1beta/models/{model}:generateContent`
- **Claude Path**: `/v1/messages`
