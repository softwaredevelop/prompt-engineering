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
	seed             = 5
	temperature      = 0.3
	topK             = 20.0
	topP             = 1
)

const (
	frequencyPenalty = 0.0
	presencePenalty  = 0.0
	thinkingBudget   = 1024 //1024, 2048, 4096, 8192
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

	relativePath := "../../../../"
	projectRoot, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("error resolving project root path: %v", err)
	}

	systemPromptFile := filepath.Join(projectRoot, "prompts", "system", "general-purpose.md")
	systemPrompt, err := gemini.ReadTextFromFile(systemPromptFile)
	if err != nil {
		log.Fatalf("error reading system instructions file: %v", err)
	}

	systemParts := []*genai.Part{
		genai.NewPartFromText(systemPrompt),
		genai.NewPartFromText("Answer questions clearly, accurately, and provide additional context when relevant."),
	}

	userPromptFile := filepath.Join(projectRoot, "prompts", "user", "hello.md")
	userPrompt, err := gemini.ReadTextFromFile(userPromptFile)
	if err != nil {
		log.Fatalf("error reading prompt instructions file: %v", err)
	}

	config := &genai.GenerateContentConfig{
		CandidateCount:    candidateCount,
		FrequencyPenalty:  gemini.F32(frequencyPenalty),
		MaxOutputTokens:   maxOutputTokens,
		PresencePenalty:   gemini.F32(presencePenalty),
		ResponseMIMEType:  responseMimeType,
		Seed:              gemini.I32(seed),
		StopSequences:     []string{"STOP!"},
		SystemInstruction: genai.NewContentFromParts(systemParts, genai.RoleUser),
		Temperature:       gemini.F32(temperature),
		TopK:              gemini.F32(topK),
		TopP:              gemini.F32(topP),
	}

	userparts := []*genai.Part{
		genai.NewPartFromText(userPrompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(userparts, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		modelName,
		contents,
		config,
	)
	if err != nil {
		log.Fatalf("failed to generate content: %v", err)
	}

	gemini.PrintResponse(result)
}
