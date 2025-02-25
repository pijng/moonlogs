package insights

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// TODO: Consider if this should be customizable
const geminiModel = "gemini-2.0-flash"

type GeminiInsightsAdapter struct {
	token string
}

func NewGeminiInsightsAdapter(token string) *GeminiInsightsAdapter {
	return &GeminiInsightsAdapter{token: token}
}

func (g *GeminiInsightsAdapter) GenerateContent(ctx context.Context, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(g.token))
	if err != nil {
		return "", fmt.Errorf("creating generative AI client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(geminiModel)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("generating content: %w", err)
	}

	cand := resp.Candidates[0]
	if cand == nil || cand.Content == nil {
		return "", errors.New("empty content candidates")
	}

	part := cand.Content.Parts[0]
	if part == nil {
		return "", errors.New("empty content parts")
	}

	return fmt.Sprint(part), nil
}
