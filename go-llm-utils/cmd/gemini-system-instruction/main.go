//revive:disable:package-comments,exported
package main

import (
	"context"
	"log"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a helpful AI assistant. Answer questions clearly and honestly.", genai.RoleUser),
	}

	contents := []*genai.Content{
		genai.NewContentFromText("Hello, what model are you?", genai.Role("user")),
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatalf("generate error: %v", err)
	}

	gemini.PrintResponse(response)
}
