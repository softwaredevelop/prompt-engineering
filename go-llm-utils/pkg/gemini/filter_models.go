package gemini

import (
	"slices"

	"google.golang.org/genai"
)

// FilterModelsByAction returns model names that support the given action.
func FilterModelsByAction(models []genai.Model, action string) []string {
	var result []string
	for _, m := range models {
		if slices.Contains(m.SupportedActions, action) {
			result = append(result, m.Name)
		}
	}
	return result
}
