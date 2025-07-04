//revive:disable:package-comments,exported
package gemini

import (
	"context"

	"google.golang.org/genai"
)

func ModelsGet(ctx context.Context, getter ModelGetter, modelName string) (*genai.Model, error) {
	return getter.Get(ctx, modelName, nil)
}
