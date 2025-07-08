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
	maxOutputTokens  = 4096         //1024, 2048, 4096, 8192, 16384
	responseMimeType = "text/plain" // text/plain, application/json, text/x.enum
	seed             = 12345
	temperature      = 0.3
	topP             = 1
)

const (
	frequencyPenalty = 0.0
)

const (
	modelName = "models/gemini-2.5-pro"
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

	systemPromptFile := filepath.Join(projectRoot, "prompts", "system", "meta-prompt.md")
	systemPrompt, err := gemini.ReadTextFromFile(systemPromptFile)
	if err != nil {
		log.Fatalf("error reading system instructions file: %v", err)
	}

	systemParts := []*genai.Part{
		genai.NewPartFromText(systemPrompt),
	}

	userPromptFile := filepath.Join(projectRoot, "prompts", "user", "prompt-generator.md")
	userPrompt, err := gemini.ReadTextFromFile(userPromptFile)
	if err != nil {
		log.Fatalf("error reading prompt instructions file: %v", err)
	}

	config := &genai.GenerateContentConfig{
		CandidateCount: candidateCount,
		// FrequencyPenalty:  gemini.F32(frequencyPenalty),
		MaxOutputTokens:   maxOutputTokens,
		ResponseMIMEType:  responseMimeType,
		Seed:              gemini.I32(seed),
		SystemInstruction: genai.NewContentFromParts(systemParts, genai.RoleUser),
		Temperature:       gemini.F32(temperature),
		TopP:              gemini.F32(topP),
	}

	userparts := []*genai.Part{
		genai.NewPartFromText(userPrompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(userparts, genai.RoleUser),
	}

	response, err := client.Models.GenerateContent(
		ctx,
		modelName,
		contents,
		config,
	)
	if err != nil {
		log.Fatalf("failed to generate content: %v", err)
	}

	responseFile := filepath.Join(projectRoot, "prompts", "user", "general-response.md")

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
