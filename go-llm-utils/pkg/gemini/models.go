//revive:disable:package-comments,exported
package gemini

import (
	"context"

	"google.golang.org/genai"
)

// ModelLister defines the interface for listing models.
type ModelLister interface {
	ListModels(ctx context.Context, config *genai.ListModelsConfig) (genai.Page[genai.Model], error)
}

// GenAIModelLister is an adapter for genai.Client.Models
type GenAIModelLister struct {
	Client *genai.Client
}

func (g *GenAIModelLister) ListModels(ctx context.Context, config *genai.ListModelsConfig) (genai.Page[genai.Model], error) {
	return g.Client.Models.List(ctx, config)
}

// ListModels returns the list of models using the provided lister.
func ListModels(ctx context.Context, lister ModelLister) (genai.Page[genai.Model], error) {
	return lister.ListModels(ctx, nil)
}

// ModelGetter defines the interface for getting a single model info.
type ModelGetter interface {
	Get(ctx context.Context, modelName string, config *genai.GetModelConfig) (*genai.Model, error)
}

type GenAIModelGetter struct {
	Client *genai.Client
}

func (g *GenAIModelGetter) Get(ctx context.Context, modelName string, config *genai.GetModelConfig) (*genai.Model, error) {
	return g.Client.Models.Get(ctx, modelName, (*genai.GetModelConfig)(config))
}
