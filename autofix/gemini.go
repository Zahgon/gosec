package autofix

import (
	"context"

	"google.golang.org/genai"
)

type GenAIModel string

const (
	ModelGeminiPro3         GenAIModel = "gemini-3-pro-preview"
	ModelGeminiPro2_5       GenAIModel = "gemini-2.5-pro"
	ModelGeminiFlash2_5     GenAIModel = "gemini-2.5-flash"
	ModelGeminiFlash2_5Lite GenAIModel = "gemini-2.5-flash-lite"
)

var _ GenAIClient = (*geminiWrapper)(nil)

type geminiWrapper struct {
	client *genai.Client
	model  GenAIModel
}

func NewGeminiClient(model, apiKey string) (GenAIClient, error) {
	_ = "STUB: not implemented"
	return *new(GenAIClient), nil
}

func (g *geminiWrapper) GenerateSolution(ctx context.Context, prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseGeminiModel(model string) (GenAIModel, error) {
	_ = "STUB: not implemented"
	return *new(GenAIModel), nil
}
