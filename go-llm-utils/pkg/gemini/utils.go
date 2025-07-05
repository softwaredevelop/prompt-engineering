//revive:disable:package-comments,exported
package gemini

import (
	"fmt"
	"os"

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
