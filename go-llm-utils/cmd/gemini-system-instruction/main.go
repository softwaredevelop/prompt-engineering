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
		SystemInstruction: genai.NewContentFromText("You are a helpful and knowledgeable AI assistant. Answer questions clearly, accurately, and provide additional context when relevant.", genai.RoleUser),
	}

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
		log.Fatalf("generate error: %v", err)
	}

	gemini.PrintResponse(response)
}
