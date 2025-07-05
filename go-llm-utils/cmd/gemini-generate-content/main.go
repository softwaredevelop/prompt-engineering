//revive:disable:package-comments,exported
package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
	"google.golang.org/genai"
)

const (
	candidateCount   = 1
	maxOutputTokens  = 8192         //1024, 2048, 4096, 8192, 16384
	responseMimeType = "text/plain" // text/plain, application/json, text/x.enum
	temperature      = 0.3
	topK             = 20.0
	topP             = 1
)

const (
	modelName = "models/gemini-2.0-flash"
)

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	relativePath := "../../../"
	projectRoot, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("Error resolving project root path: %v", err)
	}

	systemPromptPath := filepath.Join(projectRoot, "prompts", "system", "general-purpose.md")
	systemPrompt, err := gemini.ReadTextFromFile(systemPromptPath)
	if err != nil {
		log.Fatalf("Error reading system instructions file: %v", err)
	}

	userPromptPath := filepath.Join(projectRoot, "prompts", "user", "hello.md")
	userPrompt, err := gemini.ReadTextFromFile(userPromptPath)
	if err != nil {
		log.Fatalf("Error reading prompt instructions file: %v", err)
	}

	config := &genai.GenerateContentConfig{
		CandidateCount:    candidateCount,
		MaxOutputTokens:   maxOutputTokens,
		ResponseMIMEType:  responseMimeType,
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
		Temperature:       gemini.F32(temperature),
		TopK:              gemini.F32(topK),
		TopP:              gemini.F32(topP),
	}

	parts := []*genai.Part{
		genai.NewPartFromText(userPrompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		modelName,
		contents,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	gemini.PrintResponse(result)
}
