//revive:disable:package-comments,exported
package main

import (
	"context"
	"fmt"
	"log"

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

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	config := &genai.GenerateContentConfig{
		CandidateCount:   candidateCount,
		MaxOutputTokens:  maxOutputTokens,
		ResponseMIMEType: responseMimeType,
		Temperature:      gemini.F32(temperature),
		TopK:             gemini.F32(topK),
		TopP:             gemini.F32(topP),
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

	fmt.Println(response.Candidates[0].Content.Parts[0].Text)
}
