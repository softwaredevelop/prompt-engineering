//revive:disable:package-comments,exported
package gemini

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/genai"
)

func F32(v float32) *float32 { return &v }
func I32(v int32) *int32     { return &v }

func PrintResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part.Text)
			}
		}
	}
}

func ReadTextFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteGeminiTextToMarkdown(resp *genai.GenerateContentResponse, outputPath string) error {
	if resp == nil || len(resp.Candidates) == 0 {
		return fmt.Errorf("invalid or empty response from model")
	}

	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return fmt.Errorf("response candidate has no content")
	}

	var rawText string
	for _, part := range candidate.Content.Parts {
		if part.Text != "" {
			rawText += part.Text
		}
	}

	if rawText == "" {
		return fmt.Errorf("no text found in response candidate parts")
	}

	formattedText := strings.ReplaceAll(rawText, "\\n", "\n")

	err := os.WriteFile(outputPath, []byte(formattedText), 0644)
	if err != nil {
		return fmt.Errorf("failed to write markdown file %q: %w", outputPath, err)
	}

	return nil
}
