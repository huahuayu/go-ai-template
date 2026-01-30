# Go AI Template

A simple, modular Go template for calling LLM models (Gemini & Claude) via an Antigravity-compatible API gateway.

## Features

- **Modular Design**: Separated packages for `gemini` and `claude` logic.
- **Dependency Free**: Uses only Go standard library.
- **Test-Driven**: Examples provided as Go Tests for easy integration.

## Models Covered

- **Gemini**: `gemini-3-flash`, `gemini-3-pro-high`, etc.
- **Claude**: 
    - `claude-sonnet-4-5`
    - `claude-opus-4-5` (Standard Mode)
    - `claude-opus-4-5-thinking` (Thinking Mode)

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
