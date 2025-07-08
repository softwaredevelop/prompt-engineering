//revive:disable:package-comments,exported
package main

import (
	"context"
	"fmt"
	"log"
	"slices"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
)

func main() {
	ctx := context.Background()
	client, err := gemini.NewGenAIClient(ctx)
	if err != nil {
		log.Fatalf("failed to create gemini client: %v", err)
	}

	lister := &gemini.GenAIModelLister{Client: client}
	models, err := gemini.ListModels(ctx, lister)
	if err != nil {
		log.Fatalf("failed to list models: %v", err)
	}

	fmt.Println("\nList of models that support generateContent:")
	for _, model := range models.Items {
		if slices.Contains(model.SupportedActions, "generateContent") {
			fmt.Println(model.Name)
		}
	}

	fmt.Println("\nList of models that support embedContent:")
	embedModels := gemini.FilterModelsByAction(models.Items, "embedContent")
	for _, model := range embedModels {
		fmt.Println(model)
	}
}
