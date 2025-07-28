package gemini_test

import (
	"context"
	"errors"
	"testing"

	"github.com/softwaredevelop/prompt-engineering/go-llm-utils/pkg/gemini"
	"google.golang.org/genai"
)

// MockModelGetter implements the ModelGetter interface
type MockModelGetter struct {
	GetFunc func(context.Context, string, *genai.GetModelConfig) (*genai.Model, error)
}

func (m *MockModelGetter) Get(ctx context.Context, modelName string, config *genai.GetModelConfig) (*genai.Model, error) {
	return m.GetFunc(ctx, modelName, config)
}

func TestModelsGet_Success(t *testing.T) {
	mock := &MockModelGetter{
		GetFunc: func(_ context.Context, _ string, _ *genai.GetModelConfig) (*genai.Model, error) {
			return &genai.Model{
				Name:        "models/test-model",
				DisplayName: "Test Model",
				Version:     "1",
			}, nil
		},
	}

	ctx := context.Background()
	model, err := gemini.ModelsGet(ctx, mock, "models/test-model")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if model == nil {
		t.Fatal("expected model, got nil")
	}
	if model.Name != "models/test-model" {
		t.Errorf("expected model name 'models/test-model', got %s", model.Name)
	}
}

func TestModelsGet_Error(t *testing.T) {
	mock := &MockModelGetter{
		GetFunc: func(_ context.Context, _ string, _ *genai.GetModelConfig) (*genai.Model, error) {
			return nil, errors.New("model not found")
		},
	}

	ctx := context.Background()
	model, err := gemini.ModelsGet(ctx, mock, "models/unknown-model")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if model != nil {
		t.Errorf("expected nil model, got %+v", model)
	}
}
