//revive:disable:package-comments,exported
package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("failed to create gemini client: %v", err)
	}

	relativePath := "../../../../"
	projectRoot, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("failed to resolve project root path: %v", err)
	}

	responseFile := filepath.Join(projectRoot, "prompts", "user", "general-response.md")

	config := &genai.GenerateContentConfig{}

	contents := []*genai.Content{
		genai.NewContentFromText("Hello, what model are you?", genai.RoleUser),
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"models/gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatalf("failed to generate content: %v", err)
	}

	err = gemini.WriteGeminiTextToMarkdown(response, responseFile)
	if err != nil {
		log.Fatalf("failed to write response to markdown file: %v", err)
	}

	responseContent, err := gemini.ReadTextFromFile(responseFile)
	if err != nil {
		log.Fatalf("failed to read response file: %v", err)
	}

	log.Println(responseContent)
}
