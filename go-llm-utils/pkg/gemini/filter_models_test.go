package gemini

import (
	"testing"

	"google.golang.org/genai"
)

func TestFilterModelsByAction(t *testing.T) {
	models := []*genai.Model{
		{Name: "gemini-1", SupportedActions: []string{"generateContent", "embedContent"}},
		{Name: "gemini-2", SupportedActions: []string{"embedContent"}},
		{Name: "gemini-3", SupportedActions: []string{"generateContent"}},
		{Name: "gemini-4", SupportedActions: []string{}},
	}

	tests := []struct {
		action   string
		expected []string
	}{
		{"generateContent", []string{"gemini-1", "gemini-3"}},
		{"embedContent", []string{"gemini-1", "gemini-2"}},
		{"nonexistent", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.action, func(t *testing.T) {
			got := FilterModelsByAction(models, tt.action)

			if len(got) != len(tt.expected) {
				t.Fatalf("expected %d results, got %d", len(tt.expected), len(got))
			}

			for i := range tt.expected {
				if got[i] != tt.expected[i] {
					t.Errorf("at index %d: expected %q, got %q", i, tt.expected[i], got[i])
				}
			}
		})
	}
}
