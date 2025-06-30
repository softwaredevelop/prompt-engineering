package go_tools

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

// NewGenAIClient creates a new GenAI client using the GEMINI_API_KEY environment variable.
func NewGenAIClient(ctx context.Context) (*genai.Client, error) {
	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		return nil, fmt.Errorf("Environment variable GEMINI_API_KEY not set")
	}
	return newGenAIClientWithAPIKey(ctx, apiKey)
}

// newGenAIClientWithAPIKey is the internal implementation for creating a client.
// It allows injecting the API key, which is useful for testing.
func newGenAIClientWithAPIKey(ctx context.Context, apiKey string) (*genai.Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	return client, nil
}
