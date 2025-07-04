//revive:disable:package-comments,exported
package main

import (
	"context"
	"log"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
)

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	getter := &gemini.GenAIModelGetter{Client: client}
	modelName := "models/gemini-1.5-pro"

	model, err := gemini.ModelsGet(ctx, getter, modelName)
	if err != nil {
		log.Fatalf("failed to get model: %v", err)
	}

	log.Printf("Model: %s, Description: %s", model.Name, model.Description)
	log.Printf("Model: %s, Display Name: %s", model.Name, model.DisplayName)
	log.Printf("Model: %s, Input Token Limit: %d", model.Name, model.InputTokenLimit)
	log.Printf("Model: %s, Output Token Limit: %d", model.Name, model.OutputTokenLimit)
	log.Printf("Model: %s, Supported Actions: %v", model.Name, model.SupportedActions)
	log.Printf("Model: %s, Version: %s", model.Name, model.Version)
}
